package seeders

import (
	"github.com/alandwiprasetyo/rest-api/src/database"
	"github.com/alandwiprasetyo/rest-api/src/models/tables"
)

func Seeder() {
	topic1, topic2, topic3 := createTopics()
	createNews(topic1, topic2, topic3)
}

func createNews(topic tables.Topic, topic2 tables.Topic, topic3 tables.Topic) (int, int) {
	database := database.GetDatabase()
	news := tables.News{
		Headline:    "Headline 1",
		Title:       "Title 1",
		Tags:        "tag1, tag2",
		Status:      "publish",
		Description: "This is News 1 description",
		Topic: []tables.Topic{
			topic,
			topic2,
		},
	}
	database.Create(&news)

	news2 := tables.News{
		Headline:    "Headline 2",
		Title:       "Title 2",
		Tags:        "tag3, tag4",
		Status:      "draft",
		Description: "This is News 2 description",
		Topic: []tables.Topic{
			topic2,
		},
	}
	database.Create(&news2)

	news3 := tables.News{
		Headline:    "Headline 3",
		Title:       "Title 3",
		Tags:        "tag5, tag6",
		Status:      "draft",
		Description: "This is News 3 description",
		Topic: []tables.Topic{
			topic,
			topic3,
		},
	}

	database.Create(&news3)
	return news.ID, news2.ID
}

func createTopics() (tables.Topic, tables.Topic, tables.Topic) {
	database := database.GetDatabase()
	topic := tables.Topic{
		Name:        "Topic 1",
		Description: "this is description topic 1",
	}
	topic2 := tables.Topic{
		Name:        "Topic 2",
		Description: "this is description topic 2",
	}
	topic3 := tables.Topic{
		Name:        "Topic 3",
		Description: "this is description topic 3",
	}
	database.Create(&topic)
	database.Create(&topic2)
	database.Create(&topic3)
	return topic, topic2, topic3
}
