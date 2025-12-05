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
				11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
				1698522-1698528,446443-446449,38593856-38593862,565653-565659,
				824824821-824824827,2121212118-2121212124
										""";

		String sampleLine = String.join("", Arrays.asList(sampleInput.split("\n")));
		Test.assertEquals(1227775554L, solverPart1(sampleLine), "Part 1 Sample");
		Test.assertEquals(4174379265L, solverPart2(sampleLine), "Part 2 Sample");
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
			String line = readInput(inputFileName);

			// Run each of the solvers and time them
			var part1 = Timer.timeIt(() -> solverPart1(line));
			System.out.println("Part 1: " + part1.result + " ( " + Timer.formatDuration(part1.duration) + " )");

			var part2 = Timer.timeIt(() -> solverPart2(line));
			System.out.println("Part 2: " + part2.result + " ( " + Timer.formatDuration(part2.duration) + " )");

		} catch (IOException e) {
			e.printStackTrace();
			System.exit(1);
		}
	}

	// Solution to Part 1
	public static long solverPart1(String input) {
		long sum = 0;
		List<String> ids = Arrays.asList(input.split(","));
		for (String id : ids) {
			try {
				long first_id = Long.parseLong(id.split("-")[0]);
				long last_id = Long.parseLong(id.split("-")[1]);

				for (long i = first_id; i <= last_id; i++) {
					String cur_id = String.valueOf(i);
					if (cur_id.length() % 2 == 1) {
						continue;
					}
					String first_half = cur_id.substring(0, (cur_id.length() / 2));
					String last_half = cur_id.substring(cur_id.length() / 2);
					if (first_half.equals(last_half)) {
						sum += i;
					}
				}
			} catch (NumberFormatException e) {
				System.err.println("Non-numeric ids: " + id);
				return -1;
			}
		}
		return sum;
	}

	// Solution to Part 2
	public static long solverPart2(String input) {
		long sum = 0;
		List<String> ids = Arrays.asList(input.split(","));
		for (String id : ids) {
			try {
				long first_id = Long.parseLong(id.split("-")[0]);
				long last_id = Long.parseLong(id.split("-")[1]);

				for (long i = first_id; i <= last_id; i++) {
					String cur_id = String.valueOf(i);

					for (int j = 0; j <= cur_id.length() / 2; j++) {
						String r = cur_id.substring(0, j);
						if (cur_id.replaceAll(r, "").isBlank()) {
							sum += i;
							break;
						}
					}
				}
			} catch (NumberFormatException e) {
				System.err.println("Non-numeric ids: " + id);
				return -1;
			}
		}
		return sum;
	}

	// Read the input as either a file or from stdin
	private static String readInput(String filename) throws IOException {
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

		return String.join("", lines);
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
