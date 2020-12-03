import java.lang.Math;

// √(𝑥𝑝−𝑥𝑐)^2+(𝑦𝑝−𝑦𝑐)^2 < 𝑟
class Darts {
    private final double x;
    private final double y;

    public Darts(double x, double y) {
        this.x = x;
        this.y = y;
    }

    int score() {
        double location = Math.sqrt(Math.pow(this.x, 2) + Math.pow(this.y, 2));

        if (location > 10) {
            return 0;
        } else if (location <= 10 && location > 5) {
            return  1;
        } else if (location <= 5 && location > 1) {
            return 5;
        }

        return 10;
    }

}
