const Cat = artifacts.require('Cat')
const BigNumber = require('bignumber.js')
const { assertRevert } = require('../helpers/assertRevert');

contract('Cat', function(accounts){
  const symbolName = 'CAT'
  const decimals = web3.toBigNumber(18)

  const allTokens = new BigNumber(10000000000e18)

  let catInstance

  const deployer = accounts[0]
  const user1 = accounts[1]
  const user2 = accounts[2]

  beforeEach('setup contract for each test', async function () {
    catInstance = await Cat.new()
  })
describe('Initla deployment', function(){
  describe('totalSupply', function(){
    it('returns 1 * (10 ** 10) * (10 ** 18)', async function (){
      const total = await catInstance.totalSupply()

      assert.equal(total.toString(), allTokens.toString())
    })
  })

  describe('symbol', function(){
    it('returns correct symbol name', async function(){
      const symbol = await catInstance.symbol()

      assert.equal(symbol, symbolName)
    })
  })
})



describe('init balance', function () {
  describe('user1 has no tokens', function () {
    it('returns zero', async function () {
      const balance = await catInstance.balanceOf(user1)

      assert.equal(balance, 0)
    })
  })

  describe('deployer has all tokens', function () {
    it('returns the total amount of tokens', async function () {
      const balance = await catInstance.balanceOf(deployer)

      assert.equal(balance.toString(), allTokens.toString())
    })
  })
})

describe('transfer', function(){
  describe('normal transfers', async function(){
    it('transfers among deployer, user1, user2', async function(){
      // Success: transfer all tokens deployer => user1
      let amount1 = web3.toBigNumber(10000000000)
      let value1 = amount1.times(web3.toBigNumber(10).pow(decimals))

      await catInstance.transfer(user1, value1, { from: deployer })
      let deployerBalance = await catInstance.balanceOf(deployer)
      let user1Balance = await catInstance.balanceOf(user1)

      assert.equal(deployerBalance, 0)
      assert.equal(user1Balance.toString(), value1.toString())

      // Success: transfer 500 tokens user1 => user2
      let amount2 = web3.toBigNumber(500)
      let value2 = amount2.times(web3.toBigNumber(10).pow(decimals))

      await catInstance.transfer(user2, value2, { from: user1 })
      let user1Balance_new = await catInstance.balanceOf(user1)
      let user2Balance_new = await catInstance.balanceOf(user2)
      assert.equal(user2Balance_new.toString(), value2.toString())
      assert.equal(user1Balance.minus(user1Balance_new).toString(), value2.toString())
    })
  })

  describe('abnormal transfers', async function(){
    it('throws on insufficient balance', async function(){
      // Fail: insufficient funds
      let amount = web3.toBigNumber(1000)
      let value = amount.times(web3.toBigNumber(10).pow(decimals))

      await assertRevert(catInstance.transfer(user1, value, { from: user2 }))
      let user1Balance = await catInstance.balanceOf(user1)
      let user2Balance = await catInstance.balanceOf(user2)

      assert.equal(user1Balance, 0)
      assert.equal(user2Balance, 0)
    })
  })
})

describe('allow', function(){
  describe('normal allow', async function(){
    it('allows user1 to transfer from deployer', async function(){
      let amount = web3.toBigNumber(1000)
      let value = amount.times(web3.toBigNumber(10).pow(decimals))

      // Approve transfer
      await catInstance.approve(user1, value, { from: deployer })
      // Check allowance
      let allowance = await catInstance.allowance(deployer, user1, { from: user1 })
      assert.equal(allowance.toString(), value.toString())

      // Success: transferFrom
      await catInstance.transferFrom(deployer, user1, value, { from: user1 })
      // Check balance and allowance
      let allowance_new = await catInstance.allowance(deployer, user1, { from: user1 })
      assert.equal(allowance_new, 0)
      let user1Balance = await catInstance.balanceOf(user1)
      assert.equal(user1Balance.toString(), value.toString())
    })
  })

  describe('abnormal allow', async function(){
    it('throws on insufficient balance', async function(){
      let amount = web3.toBigNumber(1000)
      let value = amount.times(web3.toBigNumber(10).pow(decimals))

      // Success: approve transfer
      catInstance.approve(user2, value, { from: user1 })
      // Fail: cannot transfer
      await assertRevert(catInstance.transferFrom(user1, user2, value, { from: user1 }))
    })
  })
})

})