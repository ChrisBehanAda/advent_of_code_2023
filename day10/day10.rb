puts "Day 10:"

class Pipe
  attr_reader :from_dir, :row, :col

  def initialize(from_dir, row, col)
    @from_dir = from_dir
    @row = row
    @col = col
  end
end

def get_input(file_name)
  File.open(file_name, "r") do |file|
    file.read()
  end
end

def create_grid(input)
  input.split("\n").map { |line| line.chars}
end

def find_start(grid)
  grid.each_with_index do |row, r_idx|
    row.each_with_index do |col, c_idx|
      return [r_idx, c_idx] if grid[r_idx][c_idx] == "S"
    end
  end
  nil
end

def is_inbounds?(row, col, grid)
  row >= 0 && row < grid.length() && col >= 0 && col < grid[0].length()
end

def find_connecting_pipes(start_row, start_col, grid)
  directions = {
    "north" => [start_row - 1, start_col, '|', '7', 'F'],
    "east" =>  [start_row, start_col + 1, '-', '7', 'J'],
    "south" => [start_row + 1, start_col, '|', 'L', 'J'],
    "west" =>  [start_row, start_col - 1, '-', 'L', 'F']
  }

  pipes = []

  directions.each do |dir, ( row, col, *valid_pipes)|
    if is_inbounds?(row, col, grid) && valid_pipes.include?(grid[row][col])
      pipes << Pipe.new(get_from_direction(dir), row, col)
    end
  end

  pipes
end

def get_from_direction(dir)
  dirMap = {
    "north" => "south",
    "east" => "west",
    "south" => "north",
    "west" => "east",
  }
  return dirMap[dir]
end

def traverse_pipe(row, col, from_dir, grid)
  pipe_map = {
    "|"=> {"north"=> "south", "south"=> "north"},
    "-"=> {"east"=> "west", "west"=> "east"},
    "L"=> {"north"=> "east", "east"=> "north"},
    "J"=> {"north"=> "west", "west"=> "north"},
    "7"=> {"south"=> "west", "west"=> "south"},
    "F"=> {"south"=> "east", "east"=> "south"}
  }

  direction_change = {
    "north"=> [-1, 0],
    "east"=> [0, 1],
    "south"=> [1, 0],
    "west"=> [0, -1]
  }
  puts "row: #{row}, col: #{col}"
  pipe = grid[row][col]
  dir = pipe_map[pipe][from_dir]
  change = direction_change[dir]
  return Pipe.new(get_from_direction(dir), row + change[0], col + change[1])
end

def furthest_pipe_distance(pipe1, pipe2, grid)
  count = 1
  while true
    pipe1 = traverse_pipe(pipe1.row, pipe1.col, pipe1.from_dir, grid)
    pipe2 = traverse_pipe(pipe2.row, pipe2.col, pipe2.from_dir, grid)
    puts "Pipe 1 dir: #{pipe1.from_dir}, row: #{pipe1.row}, col: #{pipe1.col}"
    puts "Pipe 2 dir: #{pipe2.from_dir}, row: #{pipe2.row}, col: #{pipe2.col}"
    count += 1
    if pipe1.row == pipe2.row && pipe1.col == pipe2.col
      return count
    end
  end
end

def part1()
  input = get_input("input.txt")
  grid = create_grid(input)
  start_row, start_col = find_start(grid)
  pipe1, pipe2 = find_connecting_pipes(start_row, start_col, grid)
  puts "pipe1 from: #{pipe1.from_dir}"
  ans = furthest_pipe_distance(pipe1, pipe2, grid)
  puts ans
  # pipe1_dir, pipe1_row, pipe1_col = pipe1[0], pipe1[1], pipe1[2]
  # pipe2_dir, pipe2_row, pipe2_col = pipe2[0], pipe2[1], pipe2[2]
  # puts pipe1_dir, pipe1_row, pipe1_col
  # puts grid[0][2]
  # find starting position
  # Find the 2 adjacent pipes that connect to the starting position
  # traverse those pipes at the same time until the two points are the same, counting each traversal
  # return the count when the pipes are equal


end

part1()
