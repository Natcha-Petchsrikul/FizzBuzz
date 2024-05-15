import {Taxi} from './taxi'

describe('Taxi', () => {
  it('example test', () => {
    const result = Taxi(1,1)
    expect(result).toEqual(5)
  })
  it('example test2', () => {
    const result = Taxi(1.1,1)
    expect(result).toEqual(7)
  })
  it('example test3', () => {
    const result = Taxi(1.6,1)
    expect(result).toEqual(9)
  })
  it('example test4', () => {
    const result = Taxi(1.6,1.2)
    expect(result).toEqual(10)
  })
  it('example test4', () => {
    const result = Taxi(1.6,1.8)
    expect(result).toEqual(10)
  })
  it('example test4', () => {
    const result = Taxi(1,1.8)
    expect(result).toEqual(6)
  })
})
