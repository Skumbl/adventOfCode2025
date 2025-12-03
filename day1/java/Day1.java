import java.util.Scanner;
import java.util.List;
import java.util.ArrayList;
import java.io.File;
import java.io.FileNotFoundException;

public class Day1 {
    private static File file;
    private static int dialHead = 50;
    private static int zeroCounter = 0;

    public static void main(String[] args) throws FileNotFoundException {
        file = new File("input.txt");
        List<Instruction> instructions = getInstructions(file);
        for (Instruction inst : instructions) {
            int[] result = rotateDial(dialHead, inst);
            dialHead = result[0];
            zeroCounter += result[1];
        }
        System.out.println("Number of times dialHead is zero: " + zeroCounter);
    }

    private static List<Instruction> getInstructions(File file) throws FileNotFoundException {
        List<Instruction> instructions = new ArrayList<>();
        Scanner scanner = new Scanner(file);
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            char direction = line.charAt(0);
            int value = Integer.parseInt(line.substring(1));
            instructions.add(new Instruction(direction, value));
        }
        scanner.close();
        return instructions;
    }

    private static int[] rotateDial(int dialHead, Instruction inst) {
        int zeroCounter = 0;
        int valCount = inst.getValue();
        if (inst.getDirection() == 'R') {
           while (valCount > 0) {
               dialHead = Math.floorMod(dialHead + 1, 100);
               if (dialHead == 0) {
                   zeroCounter++;
               }
               valCount--;
           }

        } else {
            while (valCount > 0) {
                dialHead = Math.floorMod(dialHead - 1, 100);
                if (dialHead == 0) {
                    zeroCounter++;
                }
                valCount--;
            }
        }
        return new int[]{dialHead, zeroCounter};
    }

    static class Instruction {
        private char direction;
        private int value;
        public Instruction(char direction, int value) {
            this.direction = direction;
            this.value = value;
        }
        public char getDirection() {
            return direction;
        }
        public int getValue() {
            return value;
        }
    }
}
