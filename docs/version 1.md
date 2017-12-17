# Cell Programming Language - Version 1

A rough idea about the language. More concrete specs will follow soon.

## Operators

* `+`
* `-`
* `*`
* `/`
* `^`

## Definitions

* 
    ```clojure
        (let x 14)
    ```
* 
    ```clojure
        (fn add [x y] 
            (+ x y)
        )
    ```

## Conditionals

* 
    ```clojure
    (if (= x 14) 
        (set x 714) 
        (set x 0)
    )
    ```
* 
    ```clojure
    (cond 
        ((= x 1)
            (print "A")
        ) 
        ((= x 2)
            (print "B")
        )
    )
    ```

## Loops

* 
    ```clojure
        (loop ((let x 1)) ((< x 10))
            (
                (print x)
                (inc x)
            )
        )
    ```
