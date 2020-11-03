package com.winnerwinter.javapractice;

public class reflex_demo {
    public static void main(String[] args) throws Exception {
        String className = "com.winnerwinter.javapractice.User";
        Class<?> clazz = Class.forName(className);

        // @SuppressWarnings("unchecked")
        User user = (User) clazz.getDeclaredConstructor().newInstance();
        user.setUsername("winnerwinter");
        user.setPassword("thisismypassword");

        System.out.println(user);
    }


}

class User {
    private String username;
    private String password;

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    @Override
    public String toString() {
        return "User [username=" + username + ", password=" + password + "]";
    }
}