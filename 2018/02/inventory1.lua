#!/usr/bin/env lua

local input = arg[0]:match('(.*)/') .. '/input'

io.input(input)

local twos = 0
local threes = 0

for line in io.lines() do
  local chars = {}

  for char in line:gmatch('.') do
    if chars[char] then
      chars[char] = chars[char] + 1
    else
      chars[char] = 1
    end
  end

  for char, count in pairs(chars) do
    if count == 2 then
      twos = twos + 1
      break
    end
  end

  for char, count in pairs(chars) do
    if count == 3 then
      threes = threes + 1
      break
    end
  end
end

checksum = twos * threes

print("checksum: " .. checksum)
