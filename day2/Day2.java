import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.io.File;

// modified version of day2 to satify part 2
public class Day2 {
    public static File file = new File("input.txt");

    public static void main(String[] args) {
        List<Range> ranges = readInput(file);
        long tallyedInvalidNumbers = 0;
        for (Range range : ranges) {
            List<Long> invalidNumbers = invalidRanges(range);
            for (Long invalidNumber : invalidNumbers) {
                System.out.println("Invalid number: " + invalidNumber);
                tallyedInvalidNumbers += invalidNumber;
            }
        }
        System.out.println("Total invalid numbers: " + tallyedInvalidNumbers);
    }

    private static List<Range> readInput(File file) {
        List<Range> ranges = new ArrayList<>();
        try (Scanner scanner = new Scanner(file)) {
            scanner.useDelimiter(",");
            while (scanner.hasNext()) {
                String line = scanner.next().trim();
                String[] parts = line.split("-");
                long start = Long.parseLong(parts[0]);
                long end = Long.parseLong(parts[1]);
                Range range = new Range(start, end);
                ranges.add(range);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return ranges;
    }

    private static List<Long> invalidRanges(Range range) {
        List<Long> invalidNumbers = new ArrayList<>();
        for (long i = range.getStart(); i <= range.getEnd(); i++) {
            if (isInvalid(i)) {
                invalidNumbers.add(i);
            }
        }
        return invalidNumbers;
    }

    private static boolean isInvalid(long num) {
        String numString = Long.toString(num);
        int len = numString.length();
        for (int blocksize = 1; blocksize < len / 2 + 1; blocksize++) {
            if (len % blocksize !=0) {
                continue;
            }
            String block = numString.substring(0, blocksize);
            boolean repeating = true;
            for (int pos = blocksize; pos < len; pos += blocksize) {
                String currBlock = numString.substring(pos, pos + blocksize);
                if (!currBlock.equals(block)) {
                    repeating = false;
                    break;
                }
            }
            if (repeating) { return true;}
        }
        return false;
    }

    static class Range {
        private long start;
        private long end;

        public Range(long start, long end) {
            this.start = start;
            this.end = end;
        }
        public long getStart() {
            return start;
        }
        public long getEnd() {
            return end;
        }
    }
}
