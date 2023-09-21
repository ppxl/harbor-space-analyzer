# harbor-space-analyzer

`harbor-space-analyzer` (hsa) is a tool which polls information from a configurable Harbor instance in order to analyze size usages.

Run `hsa` with a pretty pie chart of a 25 character radius.

```bash
go run -e https://your/harbor/registry -p 8
Username: yourUsername
Password:
...
* Spongebog time card*
many debug outputs later...
...
next char A = hello-world
next char C = is-it-me-you-re-looking-for
next char D = arthur-dent
next char E = ford-prefect
next char F = slartibartfast
next char G = trillian
next char H = fenchurch
next char I = random-dent
next char J = vogon-jeltz
next char K = zaphod-beeblebrox
next char L = betelgeuse
next char M = earth
next char N = magrathea
next char O = etc
[]float64{0.0705062889105538, 0.027546107406125536, 0.0016124736165884, 0.14968270121006988, 0.002903216597359654, 0.0012146977156057162, 0.011975337218522062, 0.0004878093772923312, 0.07879950963296592, 0.4205253520611794, 0.0707865942495031, 0.0004997395875147158, 0.16194646976441593, 0.002513702652303513}
                                  
          EEEEEEFFJJJJJJ          
      EEEEEEEEEEFFJJJJJJJJKK      
    EEEEEEEEEEEEFFJJJJJJKKKKKK    
    CCEEEEEEEEEEFFJJJJKKKKKKKK    
  AACCCCEEEEEEEEFFJJJJKKKKKKKKKK  
  AAAAAACCCCEEEEFFJJKKKKKKKKKKKK  
  AAAAAAAAAACCEEFFKKKKKKKKKKKKKK  
  OOOOOOOOOOOOOOKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNNNKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNLLKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNLLKKKKKKKKKKKKKKKK  
    NNNNNNNNLLLLKKKKKKKKKKKKKK    
    NNNNNNLLLLLLKKKKKKKKKKKKKK    
      NNNNLLLLLLKKKKKKKKKKKK      
          LLLLLLKKKKKKKK        
```
