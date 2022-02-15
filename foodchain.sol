pragma solidity >= 0.4.0;

contract account {

    struct Account {
        uint id;
        string name;
        string accountAddress;
        string phoneNumber;
        uint balance;
        uint status;
        string createdTime;
    }
    Account[] public accounts;
    event Create(uint id, string createdTime);
    event Deposit(uint id, uint amount);
    event Withdraw(uint id, uint amount);
    event Transfer(uint from, uint to, uint amount);
    event Delete(uint id);
    
    function CreateAccount(uint id, string memory name, string memory accountAddress, string memory phoneNumber, uint balance, uint status, string memory createdTime) public returns (uint) {
        accounts.push(Account(id, name, accountAddress, phoneNumber, balance, status, createdTime)) ;
        emit Create(id, createdTime);
        return accounts.length;
    }
    
    function AccountInfo(uint id) view public returns(uint, string memory, string memory, string memory, uint, uint, string memory) {
        uint index = find(id);
        return (accounts[index].id, accounts[index].name, accounts[index].accountAddress, accounts[index].phoneNumber, accounts[index].balance, accounts[index].status, accounts[index].createdTime);
    }
    
    function UpdateAccount(uint id, string memory name, string memory accountAddress, string memory phoneNumber, uint balance, uint status, string memory createdTime) public returns(uint, string memory, string memory, string memory, uint, uint, string memory) {
        uint index = find(id);
        accounts[index].name = name;
        accounts[index].accountAddress = accountAddress;
        accounts[index].phoneNumber = phoneNumber;
        accounts[index].balance = balance;
        accounts[index].status = status;
        accounts[index].createdTime = createdTime;
        return (accounts[index].id, accounts[index].name, accounts[index].accountAddress, accounts[index].phoneNumber, accounts[index].balance, accounts[index].status, accounts[index].createdTime);
    }
    
    function DeleteAccount(uint id) public {
        uint index = find(id);
        accounts[index] = accounts[accounts.length - 1];
        accounts.pop();
        emit Delete(id);
    }

    function AccountTransfer(uint sender, uint receiver, uint amount) public {
        uint sID = find(sender);
        uint rID = find(receiver);

        if (amount < 0) {
            revert('please insert a valid amount !');
        }
        if (accounts[sID].balance - amount < 0) {
            revert('your balance is not enough !');
        }

        accounts[sID].balance -= amount;
        accounts[rID].balance += amount;

        emit Transfer(accounts[sID].id, accounts[rID].id, amount);

    }

    function AccountDeposit(uint id, uint amount) public {
        uint index = find(id);
        accounts[index].balance += amount;
        emit Deposit (accounts[index].id, accounts[index].balance);
    }

    function AccountWithdraw(uint id, uint amount) public {
        uint index = find(id);

        if (amount < 0) {
            revert('please insert a valid amount !');
        }
        if (accounts[index].balance - amount < 0) {
            revert('your balance is not enough !');
        }

        accounts[index].balance -= amount;
        emit Withdraw(accounts[index].id, accounts[index].balance);
    }

    
    function find(uint id) view internal returns(uint) {
        for(uint i = 0; i < accounts.length; i++) {
            if(accounts[i].id == id) {
                return i;
            }
        }
        revert('Account does not exits');
    }

    function getTotalAccounts() public view returns (uint){
        return accounts.length;
    }
}
