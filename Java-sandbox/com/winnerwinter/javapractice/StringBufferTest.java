package com.winnerwinter.javapractice;

public class StringBufferTest {
    public static void main(String[] args) {
        StringBuffer sBuffer = new StringBuffer("I love you");
        sBuffer.append("but how");
        System.out.println(sBuffer);
    }
}