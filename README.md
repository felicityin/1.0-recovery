# Run 

## Recover Privkey Slice

HBC node
```
./1.0-recovery recover -i test/hbc_input.yaml
```

Shop node
```
./1.0-recovery recover -i test/shop_input.yaml
```

## Merge Privkey Slices

```
./1.0-recovery merge -hbc 0b0d3933261ddbbce6ce4da24537aac8af78fc683532ce797c8a2d2e590a772b -shop 040e95207ece8ef6ef58d90c6fe365f3a3cf7bb7d55094c23d7e1d6e5d3549dd -chain sol
```

## Get Balance

```
# Solana
./recovery-tool balance -addr HXx8Ky1aY7GBLUghbadKais5QHdeJfdQ7mmgR9j4sqNK

# Solana USDT
./recovery-tool balance -addr HXx8Ky1aY7GBLUghbadKais5QHdeJfdQ7mmgR9j4sqNK -coin Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB
```

## Transfer

```
# Solana
./recovery-tool transfer -fromkey 078fe2333b309a95f8bc59f6e03a10c4b7b51f3e12b7ccd4a62c41363a08437a -to FUQ3cTZpuB23cohYUFWTfnK6AHTEKZ9u5vAbkBGdTFdD -amount 0.00001

# Solana USDT
./recovery-tool transfer -fromkey 078fe2333b309a95f8bc59f6e03a10c4b7b51f3e12b7ccd4a62c41363a08437a -to FUQ3cTZpuB23cohYUFWTfnK6AHTEKZ9u5vAbkBGdTFdD -amount 0.001 -coin Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB
```
