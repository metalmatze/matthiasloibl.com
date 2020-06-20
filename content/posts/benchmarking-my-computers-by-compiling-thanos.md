+++
cover = ""
date = 2020-06-20T22:00:00Z
draft = true
slug = "compiling-thanos-to-benchmark-my-computers"
title = "Compiling Thanos to Benchmark my Computers"

+++
## Why?

My desktop computer wasn't up to today's standards anymore. Although it was decent, compared to the ThinkPad X1 Carbon Gen 6 Laptop I have for work, I could feel the difference in speed on a daily basis.

Where my desktop was lacking the speed (both CPU and Disk IO), the ThinkPad was actually sometimes struggling to drive my ultra-wide monitor (DELL U3419W at 3440x1440 resolution). So, I was kind of in the unfortunate situation of having trade-offs with both (I know, pretty high standards, but I also spend 40+ hours on those computers every single week).  
Along with the COVID-19 lockdown, I figured it's the perfect time to upgrade my desktop anyway.

Thus, it was time to get a new CPU (AMD Ryzen 9 3900X 12x 3.80GHz) and SSD (Samsung 970 EVO Plus 500 GB NVMe M.2 SSD). For the new CPU I had to get a new mainboard (MSI X470) and RAM (32GB G.Skill RipJaws V DDR4-3200) too...

## How

I'm running these "benchmarks" on the [current master branch of Thanos](https://github.com/thanos-io/thanos/commit/02ab6b84908e5458699655ab064a78afc0f9da9b):  
`git checkout 02ab6b84908e5458699655ab064a78afc0f9da9b`

Each machines had the latest version of Go (1.14) installed.

To disable the Go build cache we run `go clean -cache` before each run.

```bash
go clean -cache && time go build -v ./cmd/thanos
```

## Results

### metalmatze-x1 (Thinkpad X1 Carbon 6th Gen)
```
201.17s user 15.57s system 722% cpu 29.999 total
186.64s user 14.37s system 715% cpu 28.092 total
202.74s user 15.82s system 714% cpu 30.604 total
202.26s user 15.52s system 719% cpu 30.259 total
214.77s user 17.14s system 718% cpu 32.262 total
```

Average: **30.243 seconds**
```
                       +                OS: Arch Linux x86_64
                       #                Hostname: metalmatze-x1
                      ###               Kernel Release: 5.6.15-arch1-1
                     #####              Uptime: 10 days, 10:19
                     ######             WM: i3
                    ; #####;            DE: None
                   +##.#####            Packages: 1106
                  +##########           RAM: 3764 MB / 15889 MB
                 #############;         Processor Type: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
                ###############+        $EDITOR: vim
               #######   #######        Root: 41G / 49G (83%) (ext4)
             .######;     ;###;`".      
            .#######;     ;#####.       
            #########.   .########`     
           ######'           '######    
          ;####                 ####;   
          ##'                     '##   
         #'                         `#  
```
### metalmatze-desktop (old)
```
237.25s user 18.04s system 362% cpu 1:10.33 total
237.39s user 17.99s system 361% cpu 1:10.55 total
237.92s user 18.31s system 356% cpu 1:11.82 total
238.79s user 17.93s system 363% cpu 1:10.64 total
238.44s user 17.96s system 363% cpu 1:10.58 total
```
Average: **70.784 seconds** (1:10.78)
```
                       +                OS: Arch Linux x86_64
                       #                Hostname: metalmatze-desktop
                      ###               Kernel Release: 5.5.13-arch2-1
                     #####              Uptime: 0:40
                     ######             WM: i3
                    ; #####;            DE: None
                   +##.#####            Packages: 1242
                  +##########           RAM: 1160 MB / 15996 MB
                 #############;         Processor Type: AMD FX(tm)-4100 Quad-Core Processor
                ###############+        $EDITOR: vim
               #######   #######        Root: 30G / 49G (61%) (ext4)
             .######;     ;###;`".      
            .#######;     ;#####.       
            #########.   .########`     
           ######'           '######    
          ;####                 ####;   
          ##'                     '##   
         #'                         `#  
```
### metalmatze-desktop (new)
```
136.07s user 11.17s system 1546% cpu 9.518 total
136.21s user 10.99s system 1554% cpu 9.467 total
136.74s user 10.85s system 1592% cpu 9.269 total
136.08s user 11.42s system 1550% cpu 9.514 total
136.56s user 10.89s system 1581% cpu 9.324 total
```
Average: **9.418 seconds**
```
                       +                OS: Arch Linux x86_64
                       #                Hostname: metalmatze-desktop
                      ###               Kernel Release: 5.7.4-arch1-1
                     #####              Uptime: 2:40
                     ######             WM: i3
                    ; #####;            DE: None
                   +##.#####            Packages: 512
                  +##########           RAM: 1600 MB / 32120 MB
                 #############;         Processor Type: AMD Ryzen 9 3900X 12-Core Processor
                ###############+        $EDITOR: vim
               #######   #######        Root: 7.6G / 98G (7%) (ext4)
             .######;     ;###;`".      
            .#######;     ;#####.       
            #########.   .########`     
           ######'           '######    
          ;####                 ####;   
          ##'                     '##   
         #'                         `#  
```
## Conclusion

This benchmark is far from being scientific, but tasks like these (compiling) are pretty much what I do every day, so I care a lot of the new hardware improved here. :)

Going **from 70s to 9s** is really significant and impressive. Although Go has an amazing cache for the build system, I'm pretty I'll notice the difference here and there every day going foward.

***

If you run the same benchmark, let me know on Twitter and I'll add your results down below. :)