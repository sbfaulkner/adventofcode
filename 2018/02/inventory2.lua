#!/usr/bin/env lua

local input = arg[0]:match('(.*)/') .. '/input'

io.input(input)

local ids = {}

for id in io.lines() do
  ids[#ids+1] = id
end

for i = 1, #ids do
  local iid = ids[i]
  local diff = 0

  for j = i + 1, #ids do
    local jid = ids[j]

    for c = 1, #iid do
      if iid:sub(c,c) ~= jid:sub(c,c) then
        if diff > 0 then
          diff = 0
          break
        end
        diff = c
      end
    end

    if diff > 0 then
      print("Box 1: " .. iid)
      print("Box 2: " .. jid)
      print()
      print("Common characters:" .. iid:sub(1,diff-1) .. iid:sub(diff+1,#iid))
      break
    end
  end

  if diff > 0 then
    break
  end
end
