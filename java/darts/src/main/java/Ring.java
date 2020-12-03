public class Ring {
    public double lowerBounds;
    public double upperBounds;
    public int score;

    public Ring(double upper, double lower, int score){
        this.lowerBounds = lower;
        this.upperBounds = upper;
        this.score = score;
    }

    public boolean hit(double x, double y) {
        return x <= upperBounds &&
                x >= -upperBounds &&
                y >= lowerBounds &&
                y <= -lowerBounds;
    }
}
