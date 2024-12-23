package lab04;

public class Main {

}

class SuperMan {
    private int a;

    protected SuperMan(int a) {
        this.a = a;
    }
}

// ...

class SubMan extends SuperMan {
    public SubMan(int a) {
        super(a);
    }

    public SubMan() {
        this.a = 5;
    }
}