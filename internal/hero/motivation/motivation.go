package motivation

import (
	"context"

	"github.com/meshenka/nimble/internal"
)

var motivations = []string{
	"I owe a life debt to someone in my party",
	"I owe a LOT of money to very dangerous people",
	"I need to grow in power to defeat someone who has wronged me",
	"I am searching for a way to bring a loved one from the dead",
	"I am trying to get back home",
	"I am searching for the one that stole something valuable from me",
	"I was polymorphed into another kind of creature by a Wizard",
	"Duty calls, I am honor-bound to serve",
	"I was best friend with someone",
	"I was i betrayed by someone",
	"I'm lost",
	"Wanderlust",
	"Pilgimage",
	"My hometowm is in danger",
	"Curiosity! I want to lear DEEP secrets",
	"To prove my worth",
	"I'm following a prophecy",
}

func Select(ctx context.Context) string {
	return internal.Choose(ctx, motivations)
}

func All() []string {
	return motivations
}
