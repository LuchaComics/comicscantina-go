package model

import (
    // "time"
    _ "github.com/jinzhu/gorm"
)

type Organization struct {
	ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT;"`
    Name                string `gorm:"not null; unique; size:255;"`
    Description         string `gorm:"size:1023;"`
    Email               string `gorm:"not null; unique; size:255;"`
    // Status              uint8
    OwnerID             uint64 `gorm:"index"`
    // StreetAdddress      string `gorm:"size:127;"`
    // StreetAdddressExtra string `gorm:"size:127;"`
    // City                string `gorm:"size:127;"`
    // Province            string `gorm:"size:127;"`
    // Country             string `gorm:"size:127;"`
    // Currency string `gorm:"type:varchar(127)”`
    // Language string `gorm:"type:varchar(2)”`
    // Website string `gorm:"type:varchar(255)”`
    // Phone string `gorm:"type:varchar(10)”`
    // Fax string `gorm:"type:varchar(10)”`

    // CreatedAt time.Time `gorm:"DEFAULT:current_timestamp"`
    // UpdatedAt time.Time `gorm:"DEFAULT:current_timestamp"`
    // DeletedAt time.Time

    // # Social Media
    // twitter = models.CharField(max_length=15, null=True, blank=True)
    // facebook_url = models.URLField(null=True, blank=True)
    // instagram_url = models.URLField(null=True, blank=True)
    // linkedin_url = models.URLField(null=True, blank=True)
    // github_url = models.URLField(null=True, blank=True)
    // google_url = models.URLField(null=True, blank=True)
    // youtube_url = models.URLField(null=True, blank=True)
    // flickr_url = models.URLField(null=True, blank=True)
    //
    // # Payment Processing Accounts
    // paypal_email = models.EmailField()
    //
    // # Look and Feel
    // header = models.ForeignKey(ImageUpload, null=True, blank=True, related_name='org_header',)
    // logo = models.ForeignKey(ImageUpload, null=True, blank=True, related_name='org_logo',)
    // style = models.CharField(
    //     max_length=31,
    //     choices=constants.TSHOP_THEME_OPTIONS,
    //     default='ecantina-style-5.css',
    // )
    //
    // # Users
    // administrator = models.ForeignKey(User, null=True,)
    // customers = models.ManyToManyField(Customer, blank=True)
}

// Give custom table name in our database.
func (u Organization) TableName() string {
    return "cc_organizations"
}
