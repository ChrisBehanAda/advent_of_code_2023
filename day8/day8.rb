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

  puts steps
end

part1()
