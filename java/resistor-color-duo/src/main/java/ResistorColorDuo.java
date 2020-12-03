import java.util.HashMap;

class ResistorColorDuo {
    int value(String[] colors) {
        HashMap<String, Integer> map = createMap();
        Integer first = map.get(colors[0]);
        Integer second = map.get(colors[1]);

        return Integer.parseInt(first.toString() + second.toString());
    }

    private static HashMap<String, Integer> createMap() {
        HashMap<String, Integer> map = new HashMap<String, Integer>();
        map.put("black", 0);
        map.put("brown", 1);
        map.put("red", 2);
        map.put("orange", 3);
        map.put("yellow", 4);
        map.put("green", 5);
        map.put("blue", 6);
        map.put("violet", 7);
        map.put("grey", 8);
        map.put("white", 8);

        return map;
    }
}
