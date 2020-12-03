public class DartBoard {

    private final Ring outerRing = new Ring(10, 6, 1);
    private final Ring middleRing = new Ring(5, 2, 5);
    private final Ring innerRing = new Ring(1, 0, 10);

    public int score(double x, double y) {
        if(outerRing.hit(x, y)) {
            return outerRing.score;
        }

        if (middleRing.hit(x, y)) {
            return middleRing.score;
        }

        if (innerRing.hit(x, y)) {
            return innerRing.score;
        }

        return 0;
    }
}

