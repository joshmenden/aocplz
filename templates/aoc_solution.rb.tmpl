class AOCSolution
  def initialize(filename, params: {})
    @data = File.read(filename)
    @params = params
    parse_data
  end

  def numeric?(val)
    val.match?(/\A\d+\z/)
  end

  def parse_data
  end

  def twod_arr(input)
    input.split("\n").map {|l| l.split("") }
  end

  class Node
    attr_accessor :val, :left, :right

    def initialize(val, left = nil, right = nil)
      @val = val
      @left = left
      @right = right
    end
  end

  class Graph
    attr_accessor :root

    def initialize(root = nil, raw_data = nil)
      @root = root
      @raw_data = raw_data # [[root, left, right], [root, left, right], ...etc]
      Graph.build!(@root, {}, @raw_data)
    end

    def self.build!(root, nodes, raw_data)
      nodes[root.val] = root

      root_i = raw_data.index {|rt,l,r| rt == root.val }
      return root if root_i.nil?

      rt, l, r = raw_data[root_i]
      root.left = nodes.key?(l) ? root.left = nodes[l] : root.left = Graph.build!(Node.new(l), nodes, raw_data)
      root.right = nodes.key?(r) ? root.right = nodes[r] : root.right = Graph.build!(Node.new(r), nodes, raw_data)

      return root
    end
  end

  def within_bounds?(twod_arr, r,c)
    return false if r < 0
    return false if c < 0
    return false if r >= twod_arr.size
    return false if c >= twod_arr[0].size

    return true
  end

  def extract_digits(str, single_digit: false, exclude_negative: false)
    regex = if single_digit && exclude_negative
              /\d/
            elsif single_digit
              /-?\d/
            elsif exclude_negative
              /\d+/
            else
              /-?\d+/
            end

    return str.scan(regex).map(&:to_i)
  end

  def solve!
    pt1_solution = nil
    pt1_time = Benchmark.realtime do
      pt1_solution = pt1
    end

    puts "Part 1 Solution: #{pt1_solution} (#{pt1_time.round(1)}s)"

    pt2_solution = nil
    pt2_time = Benchmark.realtime do
      pt2_solution = pt2
    end

    puts "Part 2 Solution: #{pt2_solution} (#{pt2_time.round(1)}s)"
  end
end
