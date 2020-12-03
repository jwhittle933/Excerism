import java.util.ArrayList;
import java.util.Arrays;

class ResistorColor {
    private String[] colors = {"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"};

    int colorCode(String color) {
        ArrayList<String> list = new ArrayList<String>(Arrays.asList(colors));
        return list.indexOf(color);
    }

    String[] colors() {
        return colors;
    }
}
