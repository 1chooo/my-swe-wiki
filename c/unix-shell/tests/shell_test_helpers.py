
from unittest import TestCase, TextTestResult
import re
import subprocess as proc

TIMEOUT = 30

######################
# Customize unittest #
######################

RED = "\u001b[31m"
GREEN = "\u001b[32m"
YELLOW = "\u001b[33m"
BLUE = "\u001b[34m"
MAGENTA = "\u001b[35m"
WHITE = "\u001b[37m"
RESET = "\u001b[0m"


class ShellTestCase(TestCase):
    def __init__(self, command, *args, **kwargs):
        self.shell_command = command
        super().__init__(*args, **kwargs)

    def run_shell(self, inp):
        rc, output = execute(self.shell_command, input = inp)
        self.assertEqual(rc, 0)
        return filter_shell_output(output)

class PrettierTextTestResult(TextTestResult):
    """A test result class that can print formatted text results to a stream.
    Used by TextTestRunner.
    """
    separator1 = '=' * 70
    separator2 = '-' * 70


    def __init__(self, stream, descriptions, verbosity):
        super(PrettierTextTestResult, self).__init__(stream, descriptions, verbosity)
        self.stream = stream
        self.showAll = True
        self.dots = verbosity == 1
        self.descriptions = descriptions
        self._newline = True

    def getDescription(self, test):
        doc_first_line = test.shortDescription()
        return doc_first_line

    def startTest(self, test):
        super(TextTestResult, self).startTest(test)
        if self.showAll:
            self.stream.write(self.getDescription(test))
            self.stream.write(" ... ")
            self.stream.flush()
            self._newline = False

    def _write_status(self, test, status):
        self.stream.writeln(status)
        self.stream.flush()
        self._newline = True

    def addSubTest(self, test, subtest, err):
        if err is not None:
            if self.showAll:
                if issubclass(err[0], subtest.failureException):
                    self._write_status(subtest, f"{RED}FAIL{RESET}")
                else:
                    self._write_status(subtest, f"{RED}ERROR{RESET}")
            elif self.dots:
                if issubclass(err[0], subtest.failureException):
                    self.stream.write('F')
                else:
                    self.stream.write('E')
                self.stream.flush()
        super(TextTestResult, self).addSubTest(test, subtest, err)

    def addSuccess(self, test):
        super(TextTestResult, self).addSuccess(test)
        if self.showAll:
            self._write_status(test, f"{GREEN}ok{RESET}")
        elif self.dots:
            self.stream.write('.')
            self.stream.flush()

    def addError(self, test, err):
        super(TextTestResult, self).addError(test, err)
        if self.showAll:
            self._write_status(test, f"{RED}ERROR{RESET}")
        elif self.dots:
            self.stream.write('E')
            self.stream.flush()

    def addFailure(self, test, err):
        super(TextTestResult, self).addFailure(test, err)
        if self.showAll:
            self._write_status(test, f"{RED}FAIL{RESET}")
        elif self.dots:
            self.stream.write('F')
            self.stream.flush()

    def addSkip(self, test, reason):
        super(TextTestResult, self).addSkip(test, reason)
        if self.showAll:
            self._write_status(test, "skipped {0!r}".format(reason))
        elif self.dots:
            self.stream.write("s")
            self.stream.flush()

    def addExpectedFailure(self, test, err):
        super(TextTestResult, self).addExpectedFailure(test, err)
        if self.showAll:
            self.stream.writeln("expected failure")
            self.stream.flush()
        elif self.dots:
            self.stream.write("x")
            self.stream.flush()

    def addUnexpectedSuccess(self, test):
        super(TextTestResult, self).addUnexpectedSuccess(test)
        if self.showAll:
            self.stream.writeln("unexpected success")
            self.stream.flush()
        elif self.dots:
            self.stream.write("u")
            self.stream.flush()

    def printErrors(self):
        if self.dots or self.showAll:
            self.stream.writeln()
            self.stream.flush()
        self.printErrorList('ERROR', self.errors)
        self.printErrorList('FAIL', self.failures)
        unexpectedSuccesses = getattr(self, 'unexpectedSuccesses', ())
        if unexpectedSuccesses:
            self.stream.writeln(self.separator1)
            for test in unexpectedSuccesses:
                self.stream.writeln(f"UNEXPECTED SUCCESS: {self.getDescription(test)}")
            self.stream.flush()

    def printErrorList(self, flavour, errors):
        for test, err in errors:
            self.stream.writeln(self.separator1)
            self.stream.writeln(f"%s: {MAGENTA}%s{RESET}" % (flavour,self.getDescription(test)))
            self.stream.writeln(self.separator2)
            self.stream.writeln("%s" % err)
            self.stream.flush()


####################
# Our test helpers #
####################

def sh(command):
    outb = proc.run(command, shell = True, capture_output = True).stdout
    return try_decode(outb).strip()

def execute(*args, input = None):
    if input != None:
        stdin = proc.PIPE
        input = input.encode('ASCII')
    else:
        stdin = None

    try:
        exe = proc.Popen(args, 
                         stdin = stdin, 
                         stdout = proc.PIPE, 
                         stderr = proc.STDOUT)
    except OSError as exc:
        raise RuntimeError(f"Execution Error: {exc}") from exc

    try:
        outb = exe.communicate(input = input, timeout = TIMEOUT)[0]
        out = try_decode(outb).strip()
        ret = exe.returncode
    except proc.TimeoutExpired as exc:
        try:
            exe.terminate()        # Kill Nicely
            exe.wait(timeout=.500) # Wait 500ms for exit
        except proc.TimeoutExpired:
            exe.kill()             # Process didn't exit; really kill
            exe.poll()             # Reap Defunct Process
        raise RuntimeError(f"It seems something went wrong and the program didn't finish within {TIMEOUT}s") from exc
    else:
        return (ret, out)

# inspired by https://stackoverflow.com/a/15918519
def try_decode(bytes, codecs=['ascii', 'utf8', 'latin-1']):
    exc = None
    for codec in codecs:
        try:
            return bytes.decode(codec)
        except UnicodeDecodeError as e:
            exc = e

    if exc != None:
        raise exc

def filter_shell_output(output):
    lines = output.splitlines()

    def filter_line(line):
        return re.sub(r'[Bb]ye [Bb]ye[.!]? *', '', 
                      re.sub(r'shell ?\$ *', '', 
                             re.sub(r'Welcome to mini-shell[.!]? *', '', 
                                    line)))

    def is_empty(line): return line.strip() != ''

    # Filter out parts of the shell output and remove any empty lines
    filtered_lines = filter(is_empty, map(filter_line, lines))

    return "\n".join(filtered_lines)


