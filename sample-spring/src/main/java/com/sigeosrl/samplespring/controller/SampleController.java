package com.sigeosrl.samplespring.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/sample")
@Slf4j
public class SampleController {

    @PostMapping("/test")
    public ResponseEntity<Void> test(@RequestBody String message) {
        log.info("Received message: {}", message);
        return ResponseEntity.ok().build();
    }
}
