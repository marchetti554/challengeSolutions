public class GroupAnagrams {

    public static void main(String[] args) {
        List<String> input = List.of("eat", "tea", "tan", "ate", "nat", "bat");

        Map<String, List<String>> anagramMap = new HashMap<>();

        for (String s : input) {
            String sortedString = sortString(s);
            if (anagramMap.containsKey(sortedString)) {
                anagramMap.get(sortedString).add(s);
            } else {
                anagramMap.put(sortedString, new ArrayList<>(List.of(s)));
            }
        }

        System.out.println(anagramMap.values());
    }

    private static String sortString(String s) {
        char[] stringChars = s.toCharArray();
        Arrays.sort(stringChars);
        return String.valueOf(stringChars);
    }
}
