# Card Games: Cloud-Native Card Application

This project is for learning how to build a microservices application. The application provides management of decks of cards, which provides a foundation for developing other rule based card games.

## Screenshots

Coming or not.

## Service Architecture

Card game is composed of several microservices written in Go.

Frontend - Go
Card Deck - Go
Game Rules - Go

## Data Model

Data structures map to a relational database.

User - Representing the game users information
Session - Representing the users current session
Game - Representing a game

Users can log into the system to create and participate in games. Anonymous users can read, but can't create or participate.
