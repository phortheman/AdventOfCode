import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.function.Supplier;

public class Main {
	// Run tests on the sample input to ensure it works as expected
	private static void runTests() {
		String sampleInput = """
				987654321111111
				811111111111119
				234234234234278
				818181911112111
										""";

		List<String> sampleLines = Arrays.asList(sampleInput.split("\n"));
		Test.assertEquals(357L, solverPart1(sampleLines), "Part 1 Sample");
		Test.assertEquals(3121910778619L, solverPart2(sampleLines), "Part 2 Sample");
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

	// Consolidated the logic so there is one function for both puzzles
	public static long sharedSolution(List<String> input, int digits) {
		long total_joltage = 0L;
		int max_scans = input.getFirst().length() - digits;

		for (String bank : input) {
			int scans = 0;
			int start_idx = 0;

			for (int k = digits; k > 0; k--) {
				// The first character is always assumed to be the max
				int max = Character.getNumericValue(bank.charAt(start_idx));
				start_idx++;

				// Cache the number of times we've committed to a scan
				int cur_scans = scans;

				// If we can scan, test each digit until we reach the max scans
				// Commits to a scan by updating the scans counter
				for (int i = start_idx; i < bank.length() && cur_scans < max_scans; i++) {
					cur_scans++;
					int char_val = Character.getNumericValue(bank.charAt(i));
					if (char_val > max) {
						max = char_val;
						scans = cur_scans;
						start_idx = i + 1;
					}
				}

				total_joltage += max * Math.pow(10, k - 1);
			}
		}

		return total_joltage;
	}

	// Solution to Part 1
	public static long solverPart1(List<String> input) {
		return sharedSolution(input, 2);
	}

	// Solution to Part 2
	public static long solverPart2(List<String> input) {
		return sharedSolution(input, 12);
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
