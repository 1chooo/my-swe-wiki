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
        """Recognizes a simple non-special token"""
        self.assertEqual(sh('echo a | ./tokenize'), 'a')

    def test02(self):
        """Recognizes two non-special tokens"""
        self.assertEqual(sh("echo 'a b' | ./tokenize"), "a\nb")
    
    def test03(self):
        """Recognizes three non-special multi-char tokens"""
        self.assertEqual(
                sh("echo 'foo_bar baz      hello' | ./tokenize"), 
                "foo_bar\nbaz\nhello")
  
    def test04(self):
        """Recognizes special characters as tokens"""
        self.assertEqual(sh("echo '(;|)<>' | ./tokenize"), "(\n;\n|\n)\n<\n>")

    def test05(self):
        """Recognizes a string"""
        self.assertEqual(sh("echo '\"hello world\"' | ./tokenize"), "hello world")

    def test06(self):
        """Recognizes strings mixed with other tokens"""
        self.assertEqual(
                sh("echo 'foo \"Lorem ipsum dolor sit amet\" < bar \"consectetur (adipiscing; >elit\"' | ./tokenize"), 
                "foo\nLorem ipsum dolor sit amet\n<\nbar\nconsectetur (adipiscing; >elit")
                
    def test07(self):
        """Recognizes strings with escape sequences"""
        self.assertEqual(
                sh("echo '-e \"1\\n\"' | ./tokenize"), 
                "-e\n1")

    def test08(self):
        """Recognizes strings with escape sequences"""
        self.assertEqual(
                sh("echo 'evenif\isseen newToken' | ./tokenize"), 
                'evenif\\isseen\nnewToken')      
    
    def test9(self):
        """Recognizes a simple non-special token within single quotes"""
        self.assertEqual(sh("echo 'x' | ./tokenize"), "x")

    def test10(self):
        """Recognizes multiple non-special tokens within single quotes"""
        self.assertEqual(sh("echo 'x y z' | ./tokenize"), "x\ny\nz")

    def test11(self):
        """Recognizes special characters within single quotes as tokens"""
        self.assertEqual(sh("echo '(a|b)<c>' | ./tokenize"), "(\na\n|\nb\n)\n<\nc\n>")

    def test12(self):
        """Recognizes special characters as tokens within double quotes"""
        self.assertEqual(sh('echo "(a|b)<c>" | ./tokenize'), "(\na\n|\nb\n)\n<\nc\n>")

    def test13(self):
        """Recognizes a string with escape sequences with \n and disregards next line"""
        self.assertEqual(sh("echo 'This is a newline\t abcd \n echo 'not the same input'' | ./tokenize"), "This\nis\na\nnewline\nabcd")

    def test14(self):
        """Recognizes strings with escape sequences within double quotes"""
        self.assertEqual(sh('echo "This is a \\t tab" | ./tokenize'), 'This\nis\na\ntab')

if __name__ == '__main__':
    print(f"-= {YELLOW}Running tests for {TOKENIZE}{RESET} =-")
    unittest.main(testRunner = unittest.TextTestRunner(resultclass = PrettierTextTestResult))


