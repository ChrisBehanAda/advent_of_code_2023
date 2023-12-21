puts "Day 8:"


def get_input(file_name)
  File.open(file_name, 'r') do |file|
    return file.read
  end
end


def create_map(map_lines)
  map_lines.split("\n").map do |line|
      key, vals = line.split('=').map(&:strip)
      left, right = vals[1...-1].split(',').map(&:strip)
      [key, { L: left, R: right }]
    end.to_h
end

def part1()
  input = get_input("input.txt")
  instructions, map_lines = input.split("\n\n")
  map = create_map(map_lines)

  position = "AAA"
  steps = 0

  instructions.each_char.cycle do |i|
    position = map[position][i.to_sym()]
    steps += 1
    if position == "ZZZ"
      break
    end
  end

  puts "Part 1 answer: #{steps}"
end

def get_starting_positions(positions)
  positions.select {|p| p.end_with?("A")}
end


def steps_to_end(instructions, map, position)
  steps = 0
  instructions.each_char.cycle do |i|
    position = map[position][i.to_sym()]
    steps += 1
    if position.end_with?("Z")
      return steps
    end
  end
end

def part2()
  input = get_input("input.txt")
  instructions, map_lines = input.split("\n\n")
  map = create_map(map_lines)
  positions = get_starting_positions(map.keys)
  steps = 0

  counts = positions.map {|p| steps_to_end(instructions, map, p)}

  ans = counts.reduce(:lcm)
  puts "Part 2 answer: #{ans}"
end

part1()
part2()
