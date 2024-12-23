#!/usr/bin/env python3

from unittest import TestCase

import unittest

import os.path
import sys
import subprocess
import random
import re

from shell_test_helpers import *

TOKENIZE = "./tokenize"
SHELL = "./shell"

class ShellTests(ShellTestCase):
    def __init__(self, *args, **kwargs):
        super().__init__(SHELL, *args, **kwargs)

    def test01(self):
        """ Shell prints the Welcome message and correct prompt """

        exe = subprocess.Popen(
                SHELL, 
                stdin = subprocess.DEVNULL, 
                stdout = subprocess.PIPE, 
                stderr = subprocess.STDOUT
              )

        try:
            (outb, _) = exe.communicate(timeout = 0.5)
        except subprocess.TimeoutExpired:
            exe.kill()
            (outb, _) = exe.communicate()

        out = try_decode(outb).strip()

        lines = out.splitlines()
        self.assertRegex(lines[0], "^Welcome to mini-shell.*")
        self.assertRegex(lines[1], "^shell \$.*")

    def test02(self):
        """ Exit command works """
        rc, actual = execute(SHELL, input = "exit\n")
        lines = actual.splitlines()
        matches = [re.match(".*Bye bye.", line) 
                   for line in lines[1:] 
                   if line.strip() != ""]

        self.assertTrue(
                any(matches),
                msg = "Could not find a Bye bye message")
 
    def test03(self):
        """ A single echo command works """
        output = self.run_shell("echo one")
        self.assertEqual(output, "one")

    def test04(self):
        """ Two echo commands work """
        output = self.run_shell("echo one\necho two")
        self.assertEqual(output, "one\ntwo")

    def test05(self):
        """ 'ls -a' works """
        actual = self.run_shell("ls -a")
        expected = sh("ls -a")
        self.assertEqual(actual, expected)

    def test06(self):
        """ Multiple file commands work """
        script = \
            "mkdir -p tmp\n"\
            "touch tmp/some_file\n"\
            "ls -1 tmp/\n"
        actual = self.run_shell(script)
        self.assertEqual(actual, "some_file")

    def test07(self):
        """ Multiple file commands work (without a newline at the end)"""
        script = \
            "mkdir -p tmp\n"\
            "touch tmp/some_file\n"\
            "ls -1 tmp/"
        actual = self.run_shell(script)
        self.assertEqual(actual, "some_file")

    def test08(self):
        """ Basic string tokens work """
        sh('rm -f "this is a file"')
        script = \
                'touch "this is a file"\n'\
                'ls "this is a file"'

        actual = self.run_shell(script)
        self.assertEqual(actual, "this is a file")

        sh('rm -f "this is a file"')

    def test09(self):
        """ Basic sequencing works """
        script = "echo one; echo two; echo three"
        actual = self.run_shell(script)
        self.assertEqual(actual, "one\ntwo\nthree")

    def test10(self):
        """ Piping two echo commands """
        script = "echo 'Hello' | echo 'World'"
        actual = self.run_shell(script)
        self.assertEqual(actual, "'World'")

    def test11(self):
        """ Redirecting input and output of an echo command """
        script = "echo 'Redirected input' > output.txt < input.txt"
        input_data = "This is redirected input."
        with open("input.txt", "w") as f:
            f.write(input_data)

        self.run_shell(script)

        with open("output.txt", "r") as f:
            actual = f.read().strip()

        self.assertEqual(actual, "'Redirected input'\nWelcome to mini-shell.\nshell $")

    def test12(self):
        """ Sequencing echo commands with input redirection """
        script = "echo 'Before' ; echo 'Middle' ; echo 'After' < input.txt"
        input_data = "This is redirected input."
        with open("input.txt", "w") as f:
            f.write(input_data)

        actual = self.run_shell(script)
        expected = "'Before'\n'Middle'\n'After'"
        self.assertEqual(actual, expected)

    def test13(self):
        """ Sequencing an echo command writing to a file and piping into another command """
        script = "echo 'Output to file' > output1.txt; cat output1.txt | grep 'Output'"

        self.run_shell(script)

        with open("output1.txt", "r") as f:
            actual = f.read().strip()

        self.assertEqual(actual, "'Output to file'\nWelcome to mini-shell.\nshell $")

if __name__ == '__main__':
    print(f"-= {YELLOW}Running tests for {SHELL}{RESET} =-")
    unittest.main(testRunner = unittest.TextTestRunner(resultclass = PrettierTextTestResult))


