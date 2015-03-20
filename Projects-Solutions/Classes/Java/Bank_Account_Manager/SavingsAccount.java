//Savings account
//
public class SavingsAccount extends  Account {

    //when user passes one argument
    public SavingsAccount(int accNum){
        super(accNum);
    }

    //no argument passed
    public SavingsAccount(){
        //do nothing
    }

    /*add amount to the current account balance*/
    public void deposit(double amount){
        setAccountBalance(amount + getAccountBalance());
    }

    /*subtract amount from current balance
   * throw exception if the account Balance is <= 0*/
    public double withdraw(double amount){
        if(getAccountBalance() <= amount){
            throw new RuntimeException("Insufficient funds");
        }
        setAccountBalance(getAccountBalance() - amount);

        return getAccountBalance();
    }

    /*display account balance*/
    public void displayBalance(){
        System.out.println("Account balance : " + getAccountBalance());
    }
}
