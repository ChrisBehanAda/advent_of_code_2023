puts "Day 9:"

def get_input(file_name)
  File.open(file_name, "r") do |file|
    file.read()
  end
end

def extract_histories(input)
  input.split("\n").map { |h| h.strip.split(/\s+/) }
end

def calculate_diffs(history)
  diff_list = [history]
  while !diff_list[-1].all? { |d| d == 0}
    latest_diff = diff_list[-1]
    diffs = (0...latest_diff.length() -1).map { |i| latest_diff[i+1].to_i() - latest_diff[i].to_i()}
    diff_list << diffs
  end
  return diff_list
end

def extrapolate(history)
  diff_list = calculate_diffs(history)
  next_val = 0
  sum = (diff_list.length()-1).downto(0).reduce(0) { |acc, i| acc + diff_list[i][-1].to_i()}
end

def part1()
  input = get_input("input.txt")
  histories = extract_histories(input)
  total = histories.reduce(0) {|sum, h| sum + extrapolate(h)}
  puts "Part 1 answer: #{total}"
end

def extrapolate_backwards(history)
  diff_list = calculate_diffs(history)
  diff_list[-1].unshift(0)
  (diff_list.length()-1).downto(1).each do |i|
    extrapolated_val = diff_list[i-1][0].to_i() - diff_list[i][0].to_i()
    diff_list[i-1].unshift(extrapolated_val)
  end
  return diff_list[0][0]
end

def part2()
  input = get_input("input.txt")
  histories = extract_histories(input)
  total = histories.reduce(0) { |sum, h| sum + extrapolate_backwards(h) }
  puts "Part 2 answer: #{total}"
end



part1()
part2()
