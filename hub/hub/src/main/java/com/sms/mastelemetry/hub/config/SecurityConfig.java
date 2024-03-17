package com.sms.mastelemetry.hub.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.Customizer;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.provisioning.InMemoryUserDetailsManager;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.header.writers.frameoptions.XFrameOptionsHeaderWriter;

@SuppressWarnings("deprecated") 
@Configuration
public class SecurityConfig {

    public final String LOGCLIENT = "LOGCLIENT";

    @Bean 
    public PasswordEncoder encoder() { return new BCryptPasswordEncoder(); } 

    @Bean 
    public InMemoryUserDetailsManager userDetailsService() {
        UserDetails user1 = User.builder()
            .username("client") // TODO make this dynamic read from env v
            .password(encoder().encode("password"))
            .roles(LOGCLIENT)
            .build();
        UserDetails user2 = User.builder()
            .username("client2") // TODO make this dynamic read from env v
            .password(encoder().encode("password"))
            .roles(LOGCLIENT)
            .build();
        return new InMemoryUserDetailsManager(user1, user2);
    }

    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        http
            .csrf().and()
            .cors().and()
            //.authorizeHttpRequests().anyRequest().authenticated().and()
            .authorizeHttpRequests().anyRequest().permitAll().and()
            .httpBasic(Customizer.withDefaults());
        return http.build(); 
    }



}
