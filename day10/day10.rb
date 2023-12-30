require 'set'

puts 'Day 10:'

class Pipe
  attr_reader :from_dir, :row, :col

  def initialize(from_dir, row, col)
    @from_dir = from_dir
    @row = row
    @col = col
  end
end

class Pipe2
  attr_reader :row, :col, :dirs

  def initialize(row, col, dirs)
    @row = row
    @col = col
    @dirs = dirs
  end
end

def get_input(file_name)
  File.open(file_name, 'r', &:read)
end

def create_grid(input)
  input.split("\n").map { |line| line.chars }
end

def find_start(grid)
  grid.each_with_index do |row, r_idx|
    row.each_with_index do |_col, c_idx|
      return [r_idx, c_idx] if grid[r_idx][c_idx] == 'S'
    end
  end
  nil
end

def is_inbounds?(row, col, grid)
  row >= 0 && row < grid.length && col >= 0 && col < grid[0].length
end

def get_from_direction(dir)
  dirMap = {
    'north' => 'south',
    'east' => 'west',
    'south' => 'north',
    'west' => 'east'
  }
  dirMap[dir]
end

def bfs(starting_pipe, grid)
  pipes = [starting_pipe]
  distances = grid.map do |row|
    row.map do |_col|
      -1
    end
  end
  distances[starting_pipe.row][starting_pipe.col] = 0
  while pipes.length > 0
    pipe = pipes.shift
    connected_pipes = get_connected_pipes(pipe, grid)
    pipes.concat(connected_pipes)
    connected_pipes.each do |p|
      distances[p.row][p.col] = distances[pipe.row][pipe.col] + 1 if distances[p.row][p.col] == -1
    end
  end
  distances
end

def get_connected_pipes(pipe, grid)
  pipe_map = {
    '|' => { 'north' => 'south', 'south' => 'north' },
    '-' => { 'east' => 'west', 'west' => 'east' },
    'L' => { 'north' => 'east', 'east' => 'north' },
    'J' => { 'north' => 'west', 'west' => 'north' },
    '7' => { 'south' => 'west', 'west' => 'south' },
    'F' => { 'south' => 'east', 'east' => 'south' }
  }
  directions = {
    'north' => [pipe.row - 1, pipe.col, '|', '7', 'F'],
    'east' => [pipe.row, pipe.col + 1, '-', '7', 'J'],
    'south' => [pipe.row + 1, pipe.col, '|', 'L', 'J'],
    'west' => [pipe.row, pipe.col - 1, '-', 'L', 'F']
  }

  dirs = directions.select { |dir, _| pipe.dirs.include?(dir) }

  pipes = []

  dirs.each do |dir, (r, c, *valid_pipes)|
    pipes << Pipe2.new(r, c, pipe_map[grid[r][c]][get_from_direction(dir)]) if is_inbounds?(r, c,
                                                                                            grid) && valid_pipes.include?(grid[r][c])
  end
  pipes
end

def part1
  input = get_input('day10/input.txt')
  grid = create_grid(input)
  start_row, start_col = find_start(grid)
  start = Pipe2.new(start_row, start_col, %w[north east south west])
  distances = bfs(start, grid)
  ans = distances.flatten.max
  puts "Part 1 answer: #{ans}"
end

def print_grid(grid)
  grid.each { |row| puts row.join }
end

def mark_non_loop_tiles(grid, distances)
  (0...grid.length).each do |row|
    (0...grid[0].length).each do |col|
      grid[row][col] = 'x' if distances[row][col] == -1
    end
  end
end

def is_enclosed?(grid, row, col)
  # Point in polygon algorithm which checks whether or not
  # a horizontal line going from the point to the left intersects
  # with an odd number of pipes (meaning it's enclosed in the polygon)
  return false if col == 0

  intersections = 0

  (col - 1).downto(0) do |c|
    tile = grid[row][c]
    pipes = %w[| L J]
    intersections += 1 if pipes.include?(tile) && grid[row][col] == 'x'
  end
  intersections.odd?
end

def part2
  input = get_input('test_input.txt')
  grid = create_grid(input)
  start_row, start_col = find_start(grid)
  start = Pipe2.new(start_row, start_col, %w[north east south west])
  distances = bfs(start, grid)
  mark_non_loop_tiles(grid, distances)
  print_grid(grid)
  enclosed_points = 0
  (0...grid.length).each do |row|
    (0...grid.first.length).each do |col|
      enclosed_points += 1 if is_enclosed?(grid, row, col)
    end
  end
  puts "Part 2 answer: #{enclosed_points}"
end

# part1
part2
