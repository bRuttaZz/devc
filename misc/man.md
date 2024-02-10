

// containerfiles
1. devc build <env>
	-f <dockerfile / containerfile>
	--keep-cache
2. devc rm <env>

3. devc activate env
	--no-root

// direct images
4. devc pull <image> <env>
	--no-cache

5. devc ps

6. devc rmi <image>

7. devc login registry
	
8. devc logout

9. devc prune --wipe