class Solution {
    public int evalRPN(String[] tokens) {

    	Stack<Integer> stack = new Stack<Integer>();
    	int ans = 0;

    	for (String c : tokens) {

    		switch (c) {
    			case "+": 
    				stack.push(stack.pop() + stack.pop());
    				break;
    			case "-": 
					stack.push(-(stack.pop() - stack.pop()));
    				break;
    			case "*": 
    				stack.push(stack.pop() * stack.pop());
    				break;
    			case "/": 
    				int d = stack.pop();
    				stack.push(stack.pop() / d);
    				break;    				    				    				
    			default:
    				stack.push(Integer.valueOf(c));
    		}

    	}
  		return stack.pop();
    }
}