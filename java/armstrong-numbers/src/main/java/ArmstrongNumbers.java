import java.lang.Math;

class ArmstrongNumbers {

    boolean isArmstrongNumber(int numberToCheck) {
        int numDigits = (int) Math.log10(numberToCheck) + 1;

        int sum = 0;
        int num = numberToCheck;
        while (num != 0)  {
            int rem = num % 10;
            sum += (int) Math.pow(rem, numDigits);
            num /= 10;
        }

        return sum == numberToCheck;
    }

}
