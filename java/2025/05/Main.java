import java.awt.Point;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.function.Supplier;

public class Main {
	// Run tests on the sample input to ensure it works as expected
	private static void runTests() {
		String sampleInput = """
				3-5
				10-14
				16-20
				12-18

				1
				5
				8
				11
				17
				32
				""";

		List<String> sampleLines = Arrays.asList(sampleInput.split("\n"));
		Test.assertEquals(3L, solverPart1(sampleLines), "Part 1 Sample");
		Test.assertEquals(14L, solverPart2(sampleLines), "Part 2 Sample");
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
			List<String> lines = readInput(inputFileName);

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

	// Solution to Part 1
	public static long solverPart1(List<String> input) {
		List<List<Long>> ranges = new ArrayList<>();
		int i = 0;
		long count = 0L;
		for (; i < input.size(); i++) {
			if (input.get(i).equals("")) {
				i++;
				break;
			}

			List<Long> range = new ArrayList<>();
			range.add(Long.parseLong(input.get(i).split("-")[0]));
			range.add(Long.parseLong(input.get(i).split("-")[1]));
			ranges.add(range);
		}

		for (; i < input.size(); i++) {
			Long value = Long.parseLong(input.get(i));
			for (List<Long> r : ranges) {
				if (r.get(0) <= value && value <= r.get(1)) {
					count++;
					break;
				}
			}
		}

		return count;
	}

	// Solution to Part 2
	public static long solverPart2(List<String> input) {
		List<List<Long>> ranges = new ArrayList<>();
		for (int i = 0; i < input.size(); i++) {
			if (input.get(i).equals("")) {
				i++;
				break;
			}

			List<Long> range = new ArrayList<>();
			range.add(Long.parseLong(input.get(i).split("-")[0]));
			range.add(Long.parseLong(input.get(i).split("-")[1]));
			ranges.add(range);
		}

		// Sort the ranges for easier merging
		ranges.sort((a, b) -> {
			return a.get(0).compareTo(b.get(0));
		});

		List<List<Long>> mergedRanges = new ArrayList<>();
		while (true) {
			// Last element so just add it
			if (ranges.size() == 1) {
				mergedRanges.add(ranges.get(0));
				break;
			}

			// "Pop" the first two elements to compare
			List<Long> first = ranges.remove(0);
			List<Long> second = ranges.remove(0);

			// If the first range's end is less than the seconds start they don't intersect
			if (first.get(1) < second.get(0)) {
				mergedRanges.add(first);

				// Re-add the second
				ranges.add(0, second);
			} else {
				// Create a new range
				List<Long> newRange = new ArrayList<>();
				newRange.add(first.get(0));

				// The max of the new range is the larger of the two
				newRange.add(Long.max(first.get(1), second.get(1)));

				// Add the new range to the range and re-process
				ranges.add(0, newRange);
			}

		}

		long sum = 0L;
		for (List<Long> r : mergedRanges) {
			sum += r.get(1) - r.get(0) + 1; // Handle off by one
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
