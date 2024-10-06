# Session two


## Roadmap

https://roadmap.sh/backend



## Install Linux

 After 30 years of Windows...I've switched to Linux 

I am making this post to hopefully inspire others who might be on the fence about making the transition. If you do a lot of LLM stuff, it's worth it. (I'm sure there are many thinking "duh of course it's worth it", but I hadn't seen the light until recently)

I've been slowly building up my machine by adding more graphics cards, and I take an inferencing speed hit on windows for every card I add. I want to run larger and larger models, and the overhead was getting to be too much.

Oobabooga's textgen is top notch and very efficient <3, but windows has so much overhead the inference slowdowns were becoming something I could not ignore with my current gpu setup (6x 24GB cards). There are no inferencing programs/schemes that will overcome this. I even had WSL with deepspeed installed and there was no noticeable difference in inferencing speeds compared to just windows, I tried pytorch 2.2 and there were no noticeable speed improvements in windows; this was the same for other inferencing programs too not just textgen.

https://www.reddit.com/r/Oobabooga/comments/1b1pm0m/after_30_years_of_windowsive_switched_to_linux/

## دوره جادی

الپیک ۱ - ۰۰۱ - مقدمه و شروع 

https://www.youtube.com/watch?v=cqfrsmg4BKo&list=PL7ePwBdxM4nswZ62DvL58yJZ9W4-hOLLB


## Installing Application for Golang developer After installing Ubuntu 24.04 LTS

This document outlines a comprehensive setup guide for Golang developers working on Ubuntu 22.04 LTS.
It covers a wide range of essential tools, applications, and configurations that are useful for backend development work.

https://blog.stackademic.com/installing-application-for-golang-developer-after-installing-ubuntu-24-04-lts-0444c87d7f9d


##  با «شکن» تحریم‌های اینترنتی رو بی‌اثر کن!

https://shecan.ir/



برای استفاده از شکن کافیه DNS رو تنظیم کنید
DNS: 178.22.122.100 , 185.51.200.2
DNS-over-TLS: free.shecan.ir
DNS-over-HTTPS: https://free.shecan.ir/dns-query



## Install Golang

https://go.dev/doc/install


```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

go version
```


### Make folders

```
$HOME/go/src
$HOME/go/pkg
$HOME/go/bin
```

### Add in Bashrc

```
gedit $HOME/.bashrc

```

### go version

go version go1.23.0 linux/amd64
