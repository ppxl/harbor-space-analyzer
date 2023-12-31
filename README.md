# harbor-space-analyzer

`harbor-space-analyzer` (hsa) is a tool which polls information from a configurable Harbor instance in order to analyze size usages.

Run `hsa` with a pretty pie chart of an 8 character (line, not column) radius.

```bash
go run -e analyze https://your/harbor/registry -p 8
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

You can find out what's possible with the `--help` flags on every command.

The `--pie` / `-p` let's you set the pie radius (in lines, that is)

```
go run -e analyze https://your/harbor/registry -p 16
...
                      EEEEEEEEEEFFHHJJJJJJJJ                      
                  EEEEEEEEEEEEEEFFHHJJJJJJJJJJJJ                  
              EEEEEEEEEEEEEEEEEEFFHHJJJJJJJJJJJJJJKK              
            EEEEEEEEEEEEEEEEEEEEFFHHJJJJJJJJJJJJJJKKKK            
          EEEEEEEEEEEEEEEEEEEEEEFFHHJJJJJJJJJJJJKKKKKKKK          
        EEEEEEEEEEEEEEEEEEEEEEEEFFIIJJJJJJJJJJKKKKKKKKKKKK        
      CCEEEEEEEEEEEEEEEEEEEEEEEEFFJJJJJJJJJJJJKKKKKKKKKKKKKK      
      CCCCEEEEEEEEEEEEEEEEEEEEEEFFJJJJJJJJJJKKKKKKKKKKKKKKKK      
    CCCCCCCCCCEEEEEEEEEEEEEEEEEEFFJJJJJJJJKKKKKKKKKKKKKKKKKKKK    
    AAAACCCCCCCCEEEEEEEEEEEEEEEEFFJJJJJJJJKKKKKKKKKKKKKKKKKKKK    
  AAAAAAAAAACCCCCCDDEEEEEEEEEEEEFFJJJJJJKKKKKKKKKKKKKKKKKKKKKKKK  
  AAAAAAAAAAAAAACCCCCCEEEEEEEEEEFFJJJJKKKKKKKKKKKKKKKKKKKKKKKKKK  
  AAAAAAAAAAAAAAAAAACCCCEEEEEEEEFFJJJJKKKKKKKKKKKKKKKKKKKKKKKKKK  
  AAAAAAAAAAAAAAAAAAAAAACCCCEEEEFFJJKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
  AAAAAAAAAAAAAAAAAAAAAAAAAACCEEFFKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
  OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNNNNNNNNNNNNNNNNNLLKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNNNNNNNNNNNNNNNNNLLKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNNNNNNNNNNNNNNNLLLLKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
  NNNNNNNNNNNNNNNNNNNNNNNNLLLLLLKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK  
    NNNNNNNNNNNNNNNNNNNNNNLLLLLLKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK    
    NNNNNNNNNNNNNNNNNNNNLLLLLLLLKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK    
      NNNNNNNNNNNNNNNNNNLLLLLLLLKKKKKKKKKKKKKKKKKKKKKKKKKKKK      
      NNNNNNNNNNNNNNNNLLLLLLLLLLKKKKKKKKKKKKKKKKKKKKKKKKKKKK      
        NNNNNNNNNNNNLLLLLLLLLLLLKKKKKKKKKKKKKKKKKKKKKKKKKK        
          NNNNNNNNNNLLLLLLLLLLKKKKKKKKKKKKKKKKKKKKKKKKKK          
            NNNNNNLLLLLLLLLLLLKKKKKKKKKKKKKKKKKKKKKKKK            
              NNNNLLLLLLLLLLLLKKKKKKKKKKKKKKKKKKKKKK              
                  LLLLLLLLLLLLKKKKKKKKKKKKKKKKKK                  
                      LLLLLLLLKKKKKKKKKKKKKK  
```

Also `hsa` has now color support.
