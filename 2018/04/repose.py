import os
import re
from datetime import date, datetime, timedelta
from dateutil import parser
from sys import stdout
from collections import defaultdict

INPUT = os.path.join(os.path.dirname(os.path.realpath(__file__)), "input")

RECORD_RE = re.compile(r"^\[(?P<time>[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2})\] (?P<note>.+)$")
SHIFT_RE = re.compile(r"Guard #(?P<guard>[0-9]+) begins shift")
SLEEPS = "falls asleep"
WAKES = "wakes up"

class LogEntry:
  def __init__(self, line):
    m = RECORD_RE.search(line)
    self.datetime = parser.parse(m.group('time'))
    self.note = m.group('note')
    m = SHIFT_RE.search(self.note)
    if m:
      self.guard = int(m.group('guard'))
    else:
      self.guard = None

  def __str__(self):
    return "[%s] %s" % (self.datetime, self.note)

  def __lt__(self, other):
    return self.datetime.__lt__(other.datetime)

  def isshift(self):
    return self.guard is not None

  def iswake(self):
    return self.note == WAKES

  def issleep(self):
    return self.note == SLEEPS

class Shift:
  def __init__(self, entry):
    self.guard = entry.guard
    self.start = entry.datetime
    if entry.datetime.hour > 0:
      self.start = datetime.combine(entry.datetime.date(), datetime.min.time()) + timedelta(days=1)
    self.sleeping = [False]*60

  def __str__(self):
    return "[%s] Guard #d" % (self.start, self.guard)

  def sleep(self, start):
    self.sleeping[start.minute:] = [True]*(60-start.minute)

  def wake(self, stop):
    self.sleeping[stop.minute:] = [False]*(60-stop.minute)

log = []
shifts = []
guards = defaultdict(lambda: 0)

with open(INPUT) as input:
  log = map(lambda line: LogEntry(line), input)

log.sort()

for entry in log:
  if entry.isshift():
    shifts.append(Shift(entry))
  elif entry.issleep():
    shifts[-1].sleep(entry.datetime)
  elif entry.iswake():
    shifts[-1].wake(entry.datetime)
  else:
    raise ValueError("unrecognized entry: %s" % entry)

    shifts[-1].append(entry)

print("Date   ID     Minute")
print("              000000000011111111112222222222333333333344444444445555555555")
print("              012345678901234567890123456789012345678901234567890123456789")

for shift in shifts:
  stdout.write("%02d-%02d  #%4d  " % (shift.start.month, shift.start.day, shift.guard))
  for minute in xrange(60):
    stdout.write('#' if shift.sleeping[minute] else '.')
  stdout.write("\n")
  guards[shift.guard] += shift.sleeping.count(True)

sleepiest_guard = max(guards, key=lambda key: guards[key])

sleepiest_shifts = filter(lambda shift: shift.guard == sleepiest_guard, shifts)

for shift in sleepiest_shifts:
  stdout.write("%02d-%02d  #%4d  " % (shift.start.month, shift.start.day, shift.guard))
  for minute in xrange(60):
    stdout.write('#' if shift.sleeping[minute] else '.')
  stdout.write("\n")

most = 0
sleepiest_minute = 0

for minute in xrange(60):
  minutes = map(lambda shift: shift.sleeping[minute], sleepiest_shifts).count(True)
  if minutes > most:
    sleepiest_minute = minute
    most = minutes

print("")
print("Guard # %d is sleepiest in minute %d" % (sleepiest_guard, sleepiest_minute))

sleepiest_guard = 0
sleepiest_minute = 0
frequency = 0

for guard in guards:
  guard_shifts = filter(lambda shift: shift.guard == guard, shifts)
  for minute in xrange(60):
    guard_frequency = map(lambda shift: shift.sleeping[minute], guard_shifts).count(True)
    if guard_frequency > frequency:
      sleepiest_guard = guard
      sleepiest_minute = minute
      frequency = guard_frequency

print("")
print("Guard # %d is sleepiest in minute %d (%d times)" % (sleepiest_guard, sleepiest_minute, frequency))
