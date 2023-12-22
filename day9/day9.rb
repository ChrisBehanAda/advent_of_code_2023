puts "Day 9:"

def get_input(file_name)
  File.open(file_name, "r") do |file|
    file.read()
  end
end

def extract_histories(input)
  input.split("\n").map { |h| h.strip.split(/\s+/) }
end

def extrapolate(history)
  diff_list = [history]
  while !diff_list[-1].all? { |d| d == 0}
    diffs = []
    lastest_diff = diff_list[-1]
    (0...lastest_diff.length() -1).each do |i|
      diffs << lastest_diff[i+1].to_i() - lastest_diff[i].to_i()
    end
    diff_list << diffs
  end

  next_val = 0
  (diff_list.length()-1).downto(0) do |i|
    val = diff_list[i][-1].to_i()
    # puts val
    next_val += val
  end
  return next_val
end

def part1()
  input = get_input("input.txt")
  histories = extract_histories(input)
  total = histories.reduce(0) {|sum, h| sum + extrapolate(h)}
  puts "Part 1 answer: #{total}"
end



part1()
