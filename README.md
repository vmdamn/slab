# slab

---

**slab** exists to search for two questions:
 
1. How Blockchains can bring trust to society i.e. the initial intent of having them?
2. Why the adoption among people not happening at a comparatively better pace?

*slab* is an application-specific blockchain/deterministic state-machine focused on augmenting people collaboration for web communities. It is built using CosmosSDK/Tendermint. 

> ***slab* has been released for end-user testing. You may Run App/Trigger Chain via an Instagram(IG) Handle. 
> Check out https://slab.wtf**

Contents
========

 * [Goal](#goal)
 * [Complexity and Trust](#complexity-and-trust)
 * [Halting Tokenomics](#halting-tokenomics)
 * [Slab Installation and Usage](#slab-installation-and-usage)
 * [Slab Examples](#slab-examples)
 * [Want to contribute?](#want-to-contribute)
 * [References](#references)

### Goal
---

The goal of the project is not to create another L1 blockchain. Maybe this will exist as a sidechain. We don't know but we do know that we don't want to implement a blockchain talking only about crypto coins or another novel consensus algorithm.

Not everything needs to be on Blockchain but there are enough instances in the world that are suited to be captured by Blockchains.

We are interested in pushing the Application/Problem-Domain regarding Blockchains. We think there is enough tooling out there for people to adopt blockchains and we would like to focus our attention at a higher level i.e. Given enough tooling for state-machine replication, how could people use the state-machine to solve their day-to-day problems or to capture their imagination.

The Diffusion of Innovation theory developed by Everett Rogers (1962) can be a useful model to look at here. The Diffusion of Blockchain is very low in society as compared to say Diffusion of AI. We use AI via Email Spam Detection, Long form YouTube videos/60 sec video consumption, Music/Book Recommendations, Ecommerce recommendations, etc. One way to roughly estimate the number of people who are using AI is to see the number of people who are using say YouTube: 2.5 billion active users. But how many of us use blockchain for any specific purpose in the foreground or the background? Blockchains have been around since 1990s if you consider Haber and Stornetta's work (1991).


### Complexity and Trust
---

The Slab needs to operate in a complex system environment (Think of a City, an Economy, An Ant Colony and A Living Cell as complex environments). Slab operates via a bot in an existing social-network environment (Instagram in this case). Understanding Emergent Behavior in the case of Slab is our priority. 

Any 'Trust' System (e.g. Google Reviews or Instagram/Twitter Likes/Followers) will always be probabilistic i.e. there's always a probability that can be assigned to the truth. It can be high or low or anywhere in between. However, there is still a lot of scope to improve the probabilistic system. If X people are diligent in their work and if it can be captured via Slab, those people will begin to surface among the crowd. Immutability, Traceability, and Law of large numbers should work. The same should be true for the people who will try to game the system. This kind of system is vulnerable to Sybil Attack at the end-user/user-account level where one person can create as many User Ids as possible and could try to orchestrate trust. Understanding these attacks is another goal of this project. Tools of AI could also assist here.

People have been working on the idea of Identity/Verified Credentials for quite some time now. Work has been done by World Wide Web consortium(W3C), Decentralized Identity Foundation(DIF), Trust Over IP Foundation, Rebooting the Web of Trust. We are simply curious why such work has not found adoption with the society at large! We would like to use the tool of tinkering, the scientific method and entrepreneurship to understand the situation better.


### Halting Tokenomics
---

Though Tokenomics is interesting in its own way, right now we would like to halt it. Tokens are only being used to activate the accounts on-chain because of the CosmosSDK architecture. We are considering these tokens Null and Void. We have built the Slab System to understand and capture the non-financial aspects. The Crypto Industry is ripe with rug pulls, pump and dump, etc. executed by one person/one company or a small group of people. 


### Slab Installation and Usage
---

The application-specific module that has been built for the AppChain is also named *slab*. CosmosSDK allows for the creation of custom modules that are interoperable with existing modules of the SDK. Each of these modules has its own message/transaction processor and acts as mini state-machines. IgniteCLI has been used to generate the boilerplate code.

#### Method 1: Binary

```bash
$ 
```

#### Method 2: Install From Source

```bash
$ 
```

#### Usage

```bash
$ slabd tx slab create-slab [originator-social-id] [directed-towards] [assertion] [uri-originator] [flags]
$ slabd tx slab inspect-slab [vetter-social-id] [vetting-note] [uri-vetter] [id] [flags]
$ slabd tx slab revoke-slab [revoker-social-id] [revoking-note] [uri-revoker] [id] [flags]
$ slabd q slab show-slab [id] [flags]

```



### Slab Examples
---

One way to see the *slab* is to see it as an application-specific blockchain. 
Here's another way to see the *slab*:

A *Slab* is a unit, a building block that is specific and verifiable by somebody.
This is where it differs from an 'Instagram Post' or a 'Tweet'. 

```bash

Sample 1: Our Vendor Bob has supplied 4Kgs of fabric material yesterday. His contribution as our vendor is significant.

Sample 2: Our designer Charlie has worked on this fab dress for the past 8 months. Go check it out!

Sample 3: Our designer Alice contributed an important feature that lets users use the app in night mode or day mode with the press of a button.

Sample 4: Enjoyed the Pasta and a Cappuccino made by Dan today.

```


### Want to Contribute?
---

At this stage, the best way you could contribute is by joining the Discussion on Discord. And by using Slab in weird or useful ways!

We are sorry to say that it won't be possible to add external validator nodes to the network now as we are focused on understanding the emergent behavior that the Slab would perform. Tendermint is Byzantine Fault Tolerant. So making sure that malicious nodes are not greater than a third of the nodes will take considerable effort at this point in time. Because of our use of meaningless Tokens(which were used to activate accounts on-chain), staking, slashing, etc. would be pretty absurd. Right now, We are more interested in figuring out how Slab propagates/diffuses in society and solves a specific problem while assuming that at a certain point in time validator nodes could be added to the network. Also, We are not open to feature requests at this point of time/this version.

And hey, Thanks for stopping by!


### References
---

1. Complex System : https://en.wikipedia.org/wiki/Complex_system
2. Emergent Behavior: https://www.igi-global.com/dictionary/emergent-behavior/9659
3. Emergence: https://en.wikipedia.org/wiki/Emergence
4. Haber and Stornetta: https://link.springer.com/article/10.1007/bf00196791
5. CosmosSDK Architecture: https://docs.cosmos.network/v0.46/intro/sdk-app-architecture.html

