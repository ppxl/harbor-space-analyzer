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
A:  7.05 - hello-world
C:  2.75 - is-it-me-you-re-looking-for
D:  0.16 - arthur-dent
E: 14.97 - ford-prefect
F:  0.29 - slartibartfast
G:  0.12 - trillian
H:  1.20 - fenchurch
I:  0.05 - random-dent
J:  7.88 - vogon-jeltz
K: 42.05 - zaphod-beeblebrox
L:  7.08 - betelgeuse
M:  0.05 - earth
N: 16.19 - magrathea
O:  0.15 - etc

                                  
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
