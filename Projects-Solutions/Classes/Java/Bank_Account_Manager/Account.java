/*abstract class Account*/

public abstract class Account {
    private int accountNumber;
    private double accountBalance;

    //constructor for when one argument is passed
    public Account(int accountNumber){
        this.accountNumber = accountNumber;
    }

    //constructor for when no argument is passed
    public Account(){
        //do nothing
    }

    //set accountNumber
    public void setAccountNumber(int accountNumber){
        this.accountNumber = accountNumber;
    }

    //@returns accountNumber
    public int getAccountNumber(){
        return accountNumber;
    }

    //set AccountBalance
    public void setAccountBalance(double accountBalance){
        this.accountBalance = accountBalance;
    }

    //@returns accountBalance
    public double getAccountBalance(){
        return accountBalance;
    }

    public abstract void deposit(double amount);    //deposit to the account

    public abstract void displayBalance();  //view balance

    public abstract double withdraw(double amount);   //withdraw from account
}
