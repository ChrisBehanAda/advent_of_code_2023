puts 'Day 11:'

# 1 expand universe by inserting an empty row below
# any row of all dots and an empty column to the
# right of any column with all dots.

# Assign a unique number to each galaxy in the grid
# then create a list of all unique pairs of galaxies.
# For each pair, start at the lower number in the pair and
# perform a breadth first search to find the shortest path to
# the other element in the pair.

# Or
#
# Assign a unique number to each galaxy in the grid then
# for each number perform a breadth first search that expands
# the entire grid. Each time it encounters another galaxy,
# record the shortest path to the galaxy and add it to a list of
# pairs if we have not already visited that galaxy.

def get_input(file_name)
  File.open(file_name, 'r', &:read)
end

def expand_universe(grid)
  empty_rows = []
  grid.each_with_index do |row, i|
    empty_rows << i if row.all? { |v| v == '.' }
  end

  empty_cols = []
  (0...grid.first.length).each do |col_num|
    col = grid.map { |row| row[col_num] }
    empty_cols << col_num if col.all? { |v| v == '.' }
  end

  row_insertions = 0
  empty_rows.each do |row_num|
    empty_row = grid.first.map { |_| '.' }
    grid.insert(row_num + row_insertions, empty_row)
    row_insertions += 1
  end

  col_insertions = 0
  empty_cols.each do |col_num|
    grid.each do |row|
      row.insert(col_num + col_insertions, '.')
    end
    col_insertions += 1
  end
end

def label_galaxies(universe)
  galaxy_cords = []
  count = 0
  (0...universe.length).each do |row|
    (0...universe.first.length).each do |col|
      next unless universe[row][col] == '#'

      count += 1
      universe[row][col] = count
      galaxy_cords << [row, col]
    end
  end
  galaxy_cords
end

class GalaxyPair
  attr_reader :galaxy1, :galaxy2, :shortest_distance

  def initialize(g1, g2, d)
    @galaxy1 = g1
    @galaxy2 = g2
    @shortest_distance = d
  end

  def id
    @galaxy1 <= @galaxy2 ? "#{@galaxy1}_#{@galaxy2}" : "#{@galaxy2}_#{@galaxy1}"
  end
end

def get_neighbours(universe, distances, row, col)
  dirs = [
    [-1, 0],
    [0, 1],
    [1, 0],
    [0, -1]
  ]

  neighbours = []
  dirs.each do |row_offset, col_offset|
    r = row + row_offset
    c = col + col_offset
    if r >= 0 && r < universe.length && c >= 0 && c < universe.first.length && distances[r][c] == -1
      neighbours << [r, c]
    end
  end
  neighbours
end

def galaxy_pairs(universe, row, col)
  cords = [[row, col]]
  distances = universe.map { |r| r.map { -1 } }
  distances[row][col] = 0
  g_pairs = []

  while cords.length > 0
    r, c = cords.shift
    neighbours = get_neighbours(universe, distances, r, c)
    neighbours.each do |n|
      n_r = n[0]
      n_c = n[1]
      distances[n_r][n_c] = distances[r][c] + 1
      cords << [n_r, n_c]
      if universe[n_r][n_c].is_a?(Integer)
        gp = GalaxyPair.new(universe[n_r][n_c], universe[row][col], distances[n_r][n_c])
        g_pairs << gp
      end
    end
  end
  g_pairs
end

def part1
  input = get_input('day11/input.txt')
  universe = input.split("\n").map(&:chars)
  expand_universe(universe)
  galaxies = label_galaxies(universe)
  pair_map = {}
  galaxies.each do |g|
    g_r = g[0]
    g_c = g[1]
    pairs = galaxy_pairs(universe, g_r, g_c)
    pairs.each do |p|
      pair_map[p.id] = p
    end
  end
  distance_sum = 0

  pair_map.each_value { |pair| distance_sum += pair.shortest_distance }
  puts "Part 1 answer: #{distance_sum}"
end

part1
