# Submission to SteemIt hackathon

## Inspiration

The promise of blockchains and decentralized systems is massive. Being able to provide systems that are not in control by any one entity but where participants come to consensus of how they should work is a new ability for humanity, just hitting the 9 year anniversary of Satoshi's whitepaper, and we're just getting started.

To use these systems and preserve control of their own data, users should ideally run their own full node, so they can participate in the system and avoid giving up control of their private keys.

This however creates a high barrier to entry, requiring high amounts of technical expertise to be able to configure and run e.g. a full bitcoin or steemd node.

Our prototype blockpress.me aims to solve this problem, making it easy for nontechnical users to remix content from any blockchain they want to participate in.

## What it does

Just like Wordpress is a customizable content management system that can be used by nontechnical users to create their own websites, blockpress.me aims to be an open source toolkit to let users create their own websites around blockchain data.

Anton observed that any static hosting site like Amazon S3 or IPFS can host a combination of HTML, Javascript and CSS, which pulls in data from a full blockchain node like steemd. A frontend editor can allow users to configure the layout and appearance of their site, which can pull in their posts from a steemd node, price information from a bitshares node, or balances from a bitcoind node.

If this configuration of frontend plugins is combined with a frontend+backend plugin that actually configures and runs the corresponding blockchain full nodes, users will be able to check out and host their own nodes to fully control their participation in the blockchains.

A self-service system analogous to wordpress.com could also allow users to spin up their own nodes on trusted platform providers, that only they can access.

The current prototype at blockpress.me is a proof-of-concept and a teaser for how this project could be developed.

## How we built it

The prototype was hacked up by Anton (frontend development), Dana (UI / UX) and Henrik (backend / infrastructure). 

## Challenges we ran into

We quickly set up the infrastructure and agreed on the basic idea, but some problems with the frontend development took longer than we would have liked, so Anton in particular spent a lot of SteemFest hacking on this project.

## Accomplishments that we're proud of

Taking an idea and literally nothing else and having it up on the public internet, with https and individual subdomains, nice design etc was quite inspiring.

## What we learned

We should have focused more on the concrete end-to-end features earlier on.

## What's next for blockpress.me

We're going to continue exploring this idea, to vet it against trusted advisors and the real world.
