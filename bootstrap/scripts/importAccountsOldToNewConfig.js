const { app_state: { bank: { balances } }} = require("../genesis/genesis_old.json");
const { app_state: { auth: { accounts } }} = require("../genesis/genesis_old.json");

const filteredBalances = balances.filter(({ address }) => {
  const baseAccount = accounts.find(account => account.address == address && account['@type'] === "/cosmos.auth.v1beta1.BaseAccount");
  return !!baseAccount;
}).map(({ address, coins }) => {
  const amount = BigInt(coins[0].amount);
  return { 
    address,
    amount
  };
});

const getBalance = (address) => {
  const index = filteredBalances.findIndex(balance => balance.address == address);
  return { balance: filteredBalances[index], index }
}

const changeBalance = (address, change) => {
  const { index } = getBalance(address);
  filteredBalances[index].amount += BigInt(change);
}

const getTotalBalance = () => {
  return filteredBalances.reduce((prev,cur) => prev + cur.amount, BigInt(0));
}

const wishTotalBalance = BigInt(10000000000000000);
const diffBalance = wishTotalBalance - getTotalBalance();
console.log("diffBalance:", diffBalance, "double check:", getTotalBalance() + diffBalance);

const adminAddressess = [
  "moneta1r4rfvqyjj0nycretax744krvrwlh404lp30083",
  "moneta1gyac6693yexyjzttxnr6e6qt7mfgly8fc3asf2",
  "moneta1dn92phutmq4erqze3g87lqrz4lkp20dd4pxw70",
];
// Move all assets to 0 address.
console.log(getTotalBalance());
changeBalance(adminAddressess[0], diffBalance);
adminAddressess.forEach((address) => {
  const amount = getBalance(address).balance.amount;
  changeBalance(address, BigInt(-1) * amount);
  changeBalance(adminAddressess[0], amount);
});
console.log("Total supply:", getTotalBalance());
console.log("On hands:", getTotalBalance() - getBalance("moneta1r4rfvqyjj0nycretax744krvrwlh404lp30083").balance.amount);

filteredBalances.forEach((balance) => {
  // const oldBalance = balances.find(({ address }) => address === balance.address );
  // const oldAmount = BigInt(oldBalance.coins[0].amount);
  // if (balance.amount != oldAmount) {
  //   console.log("not compare:", oldBalance, balance );
  // }

  if (balance.amount > 0n) {
    console.log(`$BIN genesis add-genesis-account ${balance.address} ${balance.amount.toString()}micromoneta`)
  }
})