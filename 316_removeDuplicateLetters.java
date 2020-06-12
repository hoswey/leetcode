class Solution {
    public String removeDuplicateLetters(String s) {

        HashMap<Character, Integer> lastPosition = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            lastPosition.put(s.charAt(i), i);
        }

        HashSet<Character> ins = new HashSet<>();
        Stack<Character> stack = new Stack<>();

        for (int i = 0; i < s.length(); i++) {

            char val = s.charAt(i) ;

            if (ins.contains(val)) {
                continue;
            }

            while (!stack.empty() && stack.peek() > val && lastPosition.get(stack.peek()) > i) {
                ins.remove(stack.pop());
            }
            
            stack.push(val);
            ins.add(val);                
            
        }

        String ans = "";
        while (!stack.empty()) {
            ans = Character.toString(stack.pop()) + ans;
        }
        return ans;
    }
}