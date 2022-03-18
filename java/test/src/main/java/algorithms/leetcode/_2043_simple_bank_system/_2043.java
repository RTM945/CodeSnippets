package algorithms.leetcode._2043_simple_bank_system;

public class _2043 {
    class Bank {

        long[] balance;

        public Bank(long[] balance) {
            this.balance = balance;
        }
        
        public boolean transfer(int account1, int account2, long money) {
            if (check(account1 - 1) && check(account2 - 1)) {
                if (balance[account1 - 1] >= money) {
                    balance[account1 - 1] -= money;
                    balance[account2 - 1] += money;
                    return true;
                }
            }
            return false;
        }
        
        public boolean deposit(int account, long money) {
            if (check(account - 1)) {
                balance[account - 1] += money;
                return true;
            }
            return false;
        }
        
        public boolean withdraw(int account, long money) {
            if (check(account - 1) && balance[account - 1] >= money) {
                balance[account - 1] -= money;
                return true;
            }
            return false;
        }

        private boolean check(int account) {
            return account >= 0 && account < balance.length;
        }
    }
}
