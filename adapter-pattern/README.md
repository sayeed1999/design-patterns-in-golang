# **Adapter Pattern in Golang**  

This sub-repository demonstrates the **Adapter Design Pattern** in Golang by integrating multiple payment providers (**Stripe, PayPal**) with different APIs into a **unified interface**.  

## **Why Should You Care?**  

Ever tried integrating **third-party APIs** and realized their method signatures **never match**?  

- One wants an **int**, another needs a **float**.  
- Some expect **currency** as a parameter, others don’t care.  
- Your clean code suddenly looks like a **patchwork of API-specific hacks**.  

Using the **Adapter Pattern**, we wrap these mismatched APIs into a **common interface**, allowing our app to consume them **seamlessly—without API headaches.**  

## **The Problem**  

You're building a checkout system and need to support **Stripe** and **PayPal**, but:  

1. **Stripe** expects `amount` as a **float** and ignores currency.  
2. **PayPal** wants `amount` as an **int** and **requires** currency.  

If we call these APIs directly, our code will be **full of conditional logic**. Instead, we use adapters to normalize them.  

## **How the Adapter Pattern Helps**  

1. **Define a Common Interface** for payments.  
2. **Create Adapters** that wrap each **third-party SDK** and implement the interface.  
3. **Let client code access payments via adapters**, without worrying about differences.  

This makes our system **modular, scalable, and easy to extend** when adding new providers. 🚀  
