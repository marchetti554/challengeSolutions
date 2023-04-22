import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class RansomNote {

    public static void main(String[] args) {
        checkMagazine(List.of("give", "me", "one", "grand", "today", "night"), List.of("give", "one", "grand", "today"));
        checkMagazine(List.of("give", "me", "one", "grand", "today", "night"), List.of("give", "one", "grand", "today"));
        checkMagazine(List.of("give", "me", "one", "grand", "today", "night"), List.of("give", "one", "grand", "today"));
    }

    public static void checkMagazine(List<String> magazine, List<String> note) {
        Map<String, Integer> magazineMap = new HashMap<>();
        for (String word : magazine) {
            magazineMap.merge(word, 1, Integer::sum);
        }
        Map<String, Integer> noteMap = new HashMap<>();
        for (String word : note) {
            noteMap.merge(word, 1, Integer::sum);
        }
        for (var word : noteMap.entrySet()) {
            if (magazineMap.get(word.getKey()) < word.getValue()) {
                System.out.println("No");
                return;
            }
        }
        System.out.println("Yes");
    }
}