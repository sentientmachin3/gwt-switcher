# Build executable
build:
    go build -o gwt-switch

# Run all go code
run:
    go run main.go git.go

# Install in ~/.local/bin the executable
install: build
    cp gwt-switch ~/.local/bin/gwt-switch

# Remove compiled executable
clean:
    rm -f gwt-switch
