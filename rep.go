package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

const (
	url = "https://www.amazon.com/-/zh/product-reviews/B086168Y25/ref=cm_cr_arp_d_paging_btm_next_%d?ie=UTF8&reviewerType=all_reviews&pageNumber=%d"
)

type Comment struct {
	Author string
	//星星
	ReviewRate string
	//评论title
	ReviewTitle string
	//评论日期
	ReviewDate string
	//内容
	ReviewContent string
	//多少人有用
	Help string
}

func main() {
	c := colly.NewCollector(
		//
		colly.MaxDepth(1),
	)
	comments := make([]*Comment, 0)
	//c1.Visit(url)
	c.OnHTML(".aok-relative", func(e *colly.HTMLElement) {
		comment := &Comment{}
		//获取作者
		comment.Author = e.ChildText("div.a-profile-content>span.a-profile-name")
		//星星
		comment.ReviewRate = e.ChildText("i.review-rating")
		//评论标题
		comment.ReviewTitle = e.ChildText("a.review-title")
		//评论日期
		comment.ReviewDate = e.ChildText("span.review-date")
		//内容
		comment.ReviewContent = e.ChildText("span.review-text-content")
		//点赞
		comment.Help = e.ChildText("span.cr-vote-text")
		comments = append(comments, comment)
	})
	for i := 1; i <=25; i++ {
		c.Visit(fmt.Sprintf(url, i, i))
	}
	for _, v := range comments {
		fmt.Println(v)
	}
}
