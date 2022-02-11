/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// addCmd represents the add command
var addEmp = &cobra.Command{
	Use:   "add",
	Short: "Adding employee name and contact in database",
	Long: `To add and create database of the employee of the oraganisation we can use
	this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		val, _ := cmd.Flags().GetString("name")
		fmt.Println(val)
		con, _ := cmd.Flags().GetString("contact")
		fmt.Println(con)
		insertData(val, con)
	},
}

func insertData(name string, contact string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:password@emp-mongo"))
	if err != nil {
		fmt.Println("Connection has not been established")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	png := client.Ping(ctx, readpref.Primary())
	fmt.Println("Tried pinging with staus", png)

	collection := client.Database("app-db").Collection("employee")
	if collection == nil {
		fmt.Println("Connection could nto be established", collection)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"Name", name}, {"Contact", contact}})
	if err != nil {
		fmt.Println("Connection has not been established", err)
		return
	}

	id := res.InsertedID
	fmt.Println(id)
}

func init() {
	rootCmd.AddCommand(addEmp)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addEmp.PersistentFlags().String("name", "", "Name of the employee")
	addEmp.MarkPersistentFlagRequired("name")
	addEmp.PersistentFlags().String("contact", "", "Contact number of the employee")
	addEmp.MarkPersistentFlagRequired("contact")

}
