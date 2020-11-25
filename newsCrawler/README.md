# How to use news crawler

## Download Golang
1. install homebrew onto your mac `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"`
2. install golang ushing homebrew `brew install go`
3. install unzip `brew install unzip`


## Download News Crawler
1. clone this repo OR update to the lastest main branch.
2. go into the newsCrawler folder and run `unzip main.zip`. You will have an "main" executable.

## Run News Crawler
1. Create a .csv file that you will store your news article URL in. `echo "" >> newsURL.csv`
2. Put article urls in _**column 1**_ 
3. Put either 0 (for fake news) **OR** 1 (for real news) in _**column 2**_ (Focus on articles only for now. No URLs from the gov websites
4. Run the following command `./main -newscsv < /path/to/your/newsURL.csv > -storecsv < /path/to/your/storeInfo.csv > -token < see WhatsApp >`




**If the output csv file looks kinda messed up, that means the website provided does not satisfy DiffBot requirements. Remove that article and find an alternative**
