package model

import (
    _ "github.com/jinzhu/gorm"
)

type Organization struct {
	ID uint64 `gorm:"primary_key";"AUTO_INCREMENT"`
    OwnerID uint64
	Name string `gorm:"type:varchar(255)”`
    Description string `gorm:"type:varchar(1023)”`
    StreetName string `gorm:"type:varchar(127)”`
    StreetNumber string `gorm:"type:varchar(127)”`
    UnitNumber string `gorm:"type:varchar(127)”`
    City string `gorm:"type:varchar(127)”`
    Province string `gorm:"type:varchar(127)”`
    Country string `gorm:"type:varchar(127)”`
    Currency string `gorm:"type:varchar(127)”`
    Language string `gorm:"type:varchar(2)”`
    Website string `gorm:"type:varchar(255)”`
    Phone string `gorm:"type:varchar(10)”`
    Fax string `gorm:"type:varchar(10)”`
    Email string `gorm:"type:varchar(127)”`

    // joined = models.DateTimeField(auto_now_add=True)
    // last_updated = models.DateTimeField(auto_now=True)
    //
    // # Variable controls whether the Organization is no longer listed in our
    // # system and Users are not allowed to login/access it.
    // is_suspended = models.BooleanField(default=False, db_index=True)
    //
    // # Variable controls whether we are to allow displaying and listing
    // # this Organization in our system. Setting to "False" means it won't
    // # appear anywhere. This value is read-only and is only adjusted
    // # by the staff of eCantina to set it 'False'.
    // is_listed = models.BooleanField(default=True, db_index=True)
    //
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
