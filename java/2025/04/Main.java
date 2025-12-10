import java.awt.Point;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.function.Supplier;
import java.util.stream.Collectors;

public class Main {
	// Run tests on the sample input to ensure it works as expected
	private static void runTests() {
		String input = """
						..@@.@@@@.
						@@@.@.@.@@
						@@@@@.@.@@
						@.@@@@..@.
						@@.@@@@.@@
						.@@@@@@@.@
						.@.@.@.@@@
						@.@@@.@@@@
						.@@@@@@@@.
						@.@.@@@.@.
				""";
		List<String> sampleInput = Arrays.stream(input.split("\n"))
				.map(String::trim)
				.collect(Collectors.toList());

		List<List<String>> sample2DArr = convertInput(sampleInput);

		Test.assertEquals(13L, solverPart1(sample2DArr), "Part 1 Sample");
		Test.assertEquals(43L, solverPart2(sample2DArr), "Part 2 Sample");
	}

	public static void main(String[] args) {
		// Test the sample input
		runTests();

		// Read the arguments
		String inputFileName = null;
		for (int i = 0; i < args.length; i++) {
			switch (args[i]) {
				case "-f":
				case "--file":
					if (i + 1 < args.length) {
						inputFileName = args[++i];
					}
					break;
				default:
					System.err.println("No arguments provided");
					System.exit(1);
			}
		}

		try {
			// Read the inputs
			List<List<String>> lines = convertInput(readInput(inputFileName));

			// Run each of the solvers and time them
			var part1 = Timer.timeIt(() -> solverPart1(lines));
			System.out.println("Part 1: " + part1.result + " ( " + Timer.formatDuration(part1.duration) + " )");

			var part2 = Timer.timeIt(() -> solverPart2(lines));
			System.out.println("Part 2: " + part2.result + " ( " + Timer.formatDuration(part2.duration) + " )");

		} catch (IOException e) {
			e.printStackTrace();
			System.exit(1);
		}
	}

	// Check if the point is in range and is a roll of paper
	private static long valid(int x, int y, List<List<String>> input) {
		if (x < 0 || y < 0 || y >= input.size() || x >= input.get(y).size()) {
			return 0L;
		}
		return input.get(y).get(x).equals("@") ? 1L : 0L;
	}

	// Check around the point for rolls of paper, returns true if less than 4 exist
	private static boolean scan(int x, int y, List<List<String>> input) {
		long count = 0L;

		if (!input.get(y).get(x).equals("@")) {
			return false;
		}

		// North
		count += valid(x - 1, y - 1, input);
		count += valid(x, y - 1, input);
		count += valid(x + 1, y - 1, input);

		// Middle
		count += valid(x - 1, y, input);
		count += valid(x + 1, y, input);

		// South
		count += valid(x - 1, y + 1, input);
		count += valid(x, y + 1, input);
		count += valid(x + 1, y + 1, input);

		return count < 4;
	}

	// Solution to Part 1
	public static long solverPart1(List<List<String>> input) {
		long sum = 0L;
		for (int y = 0; y < input.size(); y++) {
			for (int x = 0; x < input.get(y).size(); x++) {
				if (scan(x, y, input)) {
					sum++;
				}
			}
		}
		return sum;
	}

	// Solution to Part 2
	public static long solverPart2(List<List<String>> input) {
		long sum = 0L;
		long prev = -1L;
		List<Point> cache = new ArrayList<>();

		// If the previous sum is the same as the current then
		// no new valid towels were found
		while (sum != prev) {
			if (sum != prev) {
				prev = sum;
			}

			for (int y = 0; y < input.size(); y++) {
				for (int x = 0; x < input.get(y).size(); x++) {
					Point p = new Point(x, y);
					if (scan(x, y, input)) {
						// Cache the point for later updates
						cache.add(p);
						sum++;
					}
				}
			}

			// Update the input so that we remove already removed towels
			for (Point p : cache) {
				input.get(p.y).set(p.x, "x");
			}

			// Reset the list and start over
			cache.clear();
		}

		return sum;
	}

	// Read the input as either a file or from stdin
	private static List<String> readInput(String filename) throws IOException {
		BufferedReader reader;
		if (filename == null) {
			reader = new BufferedReader(new InputStreamReader(System.in));
		} else {
			reader = new BufferedReader(new FileReader(filename));
		}

		List<String> lines = new ArrayList<>();
		String line;
		while ((line = reader.readLine()) != null) {
			lines.add(line);
		}
		reader.close();

		return lines;
	}

	// Makes it easier to convert the example
	private static List<List<String>> convertInput(List<String> input) {
		List<List<String>> arr = new ArrayList<>(input.size());
		for (String line : input) {
			arr.add(Arrays.asList(line.split("")));
		}
		return arr;
	}

}

// Basic unit test class for testing sample input
class Test {
	public static <T> void assertEquals(T expected, T actual, String name) {
		if (!expected.equals(actual)) {
			System.err.println("FAIL: " + name);
			System.err.println("  Expected: " + expected);
			System.err.println("  Actual: " + actual);
			System.exit(1);
		}
	}
}

// Timer utility class for testing efficiency
class Timer {
	static class Result<T> {
		final T result;
		final long duration;

		Result(T result, long nanos) {
			this.result = result;
			this.duration = nanos;
		}
	}

	static <T> Result<T> timeIt(Supplier<T> fn) {
		long start = System.nanoTime();
		T result = fn.get();
		long duration = System.nanoTime() - start;
		return new Result<>(result, duration);
	}

	static String formatDuration(long duration) {
		double seconds = duration / 1_000_000_000.0;
		double millis = duration / 1_000_000.0;
		double micros = duration / 1_000.0;

		if (seconds >= 1.0) {
			return String.format("%.2f s", seconds);
		} else if (millis >= 1.0) {
			return String.format("%.1f ms", millis);
		} else if (micros >= 1.0) {
			return String.format("%.1f μs", micros);
		}
		return String.format("%d ns", duration);
	}
}
